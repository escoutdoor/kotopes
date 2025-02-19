FROM golang:apline3.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bin/auth ./cmd/grcp_server/main.go

FROM alpine:3.20
WORKDIR /app/
COPY --from=builder /app/bin/auth .

CMD ["/app/auth"]
