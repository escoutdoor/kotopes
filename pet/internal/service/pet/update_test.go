package pet

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (s *ServiceSuite) TestUpdate_Success() {
	var (
		id          = gofakeit.UUID()
		ownerID     = gofakeit.UUID()
		name        = gofakeit.Name()
		description = gofakeit.Paragraph(1, 1, 1, "")
		age         = int32(gofakeit.IntRange(1, 20))
	)
	pet := &model.Pet{
		ID:      id,
		OwnerID: ownerID,
	}

	in := &model.UpdatePet{
		ID:          id,
		OwnerID:     ownerID,
		Name:        &name,
		Description: &description,
		Age:         &age,
	}

	s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)
	s.txManager.ReadCommittedMock.Set(
		func(ctx context.Context, fn func(ctx context.Context) error) error {
			return fn(ctx)
		},
	)
	s.petRepo.UpdateMock.Expect(s.ctx, in).Return(nil)
	s.petCache.DeleteMock.Expect(s.ctx, id).Return(nil)

	err := s.service.Update(s.ctx, in)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestUpdate_Error() {
	var (
		id          = gofakeit.UUID()
		ownerID     = gofakeit.UUID()
		name        = gofakeit.Name()
		description = gofakeit.Paragraph(1, 1, 1, "")
		age         = int32(gofakeit.IntRange(1, 20))

		repoGetByIDErr = fmt.Errorf("repository error: could not get pet by id")
		repoUpdateErr  = fmt.Errorf("repository error: could not update pet")
		cacheErr       = fmt.Errorf("delete cache error")
	)
	pet := &model.Pet{
		ID:      id,
		OwnerID: ownerID,
	}

	in := &model.UpdatePet{
		ID:          id,
		OwnerID:     ownerID,
		Name:        &name,
		Description: &description,
		Age:         &age,
	}

	tests := []struct {
		name        string
		expectedErr error
		on          func()
	}{
		{
			name:        "repository_error_get_by_id",
			expectedErr: repoGetByIDErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(nil, repoGetByIDErr)
			},
		},
		{
			name:        "error_not_pet_owner",
			expectedErr: model.ErrNotPetOwner,
			on: func() {
				s.petRepo.GetByIDMock.
					Expect(s.ctx, id).
					Return(&model.Pet{OwnerID: gofakeit.UUID()}, nil)
			},
		},
		{
			name:        "repository_error_update",
			expectedErr: repoUpdateErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)
				s.txManager.ReadCommittedMock.Set(
					func(ctx context.Context, fn func(ctx context.Context) error) error {
						return fn(ctx)
					},
				)

				s.petRepo.UpdateMock.Expect(s.ctx, in).Return(repoUpdateErr)
			},
		},
		{
			name:        "cache_error_delete",
			expectedErr: cacheErr,
			on: func() {
				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)
				s.txManager.ReadCommittedMock.Set(
					func(ctx context.Context, fn func(ctx context.Context) error) error {
						return fn(ctx)
					},
				)
				s.petRepo.UpdateMock.Expect(s.ctx, in).Return(nil)

				s.petCache.DeleteMock.Expect(s.ctx, id).Return(cacheErr)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()
			tt.on()

			err := s.service.Update(s.ctx, in)
			s.Require().Error(err)
			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}
