package httputil

import (
	"errors"
	"net/http"

	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/go-chi/render"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string                 `json:"status"`          // user-level status message
	AppCode    int64                  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string                 `json:"error,omitempty"` // application-level error message, for debugging
	Errors     []ValidationErrorField `json:"errors,omitempty"`
}

type ValidationErrorField struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrValidation(errs []ValidationErrorField) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad Request",
		Errors:         errs,
	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad Request",
		ErrorText:      err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     "Unauthorized",
		ErrorText:      err.Error(),
	}
}

func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		StatusText:     "Forbidden",
		ErrorText:      err.Error(),
	}
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     "Not Found",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server",
		ErrorText:      err.Error(),
	}
}

func HandleGrpcError(w http.ResponseWriter, r *http.Request, err error) {
	st, ok := status.FromError(err)
	if !ok {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	logger.Infof(r.Context(), "st error: %s\n", st.Err())
	logger.Infof(r.Context(), "st message: %s\n", st.Message())
	logger.Infof(r.Context(), "st string: %s\n", st.String())

	switch st.Code() {
	case codes.InvalidArgument:
		var validationErrors []ValidationErrorField
		for _, detail := range st.Details() {
			if badRequest, ok := detail.(*errdetails.BadRequest); ok {
				for _, violation := range badRequest.GetFieldViolations() {
					validationErrors = append(validationErrors, ValidationErrorField{
						Field:       violation.GetField(),
						Description: violation.GetDescription(),
					})
				}
			}
		}

		if len(validationErrors) > 0 {
			render.Render(w, r, ErrValidation(validationErrors))
			return
		}

		render.Render(w, r, ErrBadRequest(errors.New(st.Message())))
	case codes.NotFound:
		render.Render(w, r, &ErrResponse{
			HTTPStatusCode: http.StatusNotFound,
			StatusText:     "Not Found",
			ErrorText:      st.Message(),
		})
	case codes.Unauthenticated:
		render.Render(w, r, ErrUnauthorized(errors.New(st.Message())))
	case codes.PermissionDenied:
		render.Render(w, r, ErrForbidden(errors.New(st.Message())))
	case codes.AlreadyExists:
		render.Render(w, r, ErrBadRequest(errors.New(st.Message())))
	default:
		// render.Render(w, r, ErrInternalServer(errors.New("internal server error")))
		render.Render(w, r, ErrInternalServer(err))
	}
}
