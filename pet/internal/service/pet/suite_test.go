package pet

import (
	"context"
	"testing"
	"time"

	dbmocks "github.com/escoutdoor/kotopes/common/pkg/mocks"
	repomocks "github.com/escoutdoor/kotopes/pet/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	txManager *dbmocks.TxManagerMock
	petRepo   *repomocks.PetRepositoryMock
	petCache  *repomocks.PetCacheMock
	ttl       time.Duration

	ctx     context.Context
	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.txManager = dbmocks.NewTxManagerMock(s.T())
	s.petRepo = repomocks.NewPetRepositoryMock(s.T())
	s.petCache = repomocks.NewPetCacheMock(s.T())
	s.ttl = time.Minute * 5

	s.service = NewService(
		s.petRepo,
		s.petCache,
		s.txManager,
		s.ttl,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestPetServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
