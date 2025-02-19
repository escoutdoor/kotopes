package pet

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (s *ServiceSuite) TestDelete_Success() {
	var (
		id     = gofakeit.UUID()
		userID = gofakeit.UUID()
	)

	pet := &model.Pet{
		ID:      id,
		OwnerID: userID,
	}

	s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)

	s.txManager.ReadCommittedMock.Set(
		func(ctx context.Context, fn func(ctx context.Context) error) error {
			return fn(ctx)
		},
	)
	s.petRepo.DeleteMock.Expect(s.ctx, id).Return(nil)
	s.petCache.DeleteMock.Expect(s.ctx, id).Return(nil)

	err := s.service.Delete(s.ctx, &model.DeletePet{ID: id, OwnerID: userID})
	s.Require().Nil(err)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestDelete_Error() {
	var (
		id     = gofakeit.UUID()
		userID = gofakeit.UUID()

		repoErr        = errors.New("repository error")
		redisClientErr = errors.New("redis client error")
	)

	pet := &model.Pet{
		ID:      id,
		OwnerID: userID,
	}

	tests := []struct {
		name        string
		expectedErr error
		on          func()
	}{
		{
			name:        "repository_error_get_by_id",
			expectedErr: repoErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(nil, repoErr)
			},
		},
		{
			name:        "repository_error_delete",
			expectedErr: repoErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)

				s.txManager.ReadCommittedMock.Set(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
				s.petRepo.DeleteMock.Expect(s.ctx, id).Return(repoErr)
			},
		},
		{
			name:        "cache_error_delete",
			expectedErr: redisClientErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)

				s.txManager.ReadCommittedMock.Set(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
				s.petRepo.DeleteMock.Expect(s.ctx, id).Return(nil)
				s.petCache.DeleteMock.Expect(s.ctx, id).Return(redisClientErr)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()

			tt.on()

			err := s.service.Delete(s.ctx, &model.DeletePet{ID: id, OwnerID: userID})
			s.Require().Error(err)
			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}
