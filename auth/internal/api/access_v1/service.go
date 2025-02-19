package access_v1

import (
	"github.com/escoutdoor/kotopes/auth/internal/service"
	pb "github.com/escoutdoor/kotopes/common/api/access/v1"
)

type Implementation struct {
	pb.UnimplementedAccessV1Server

	accessService service.AccessService
}

func New(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
