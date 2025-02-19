package pet

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (s *ServiceSuite) TestGet_Success() {
	var (
		id          = gofakeit.UUID()
		petCacheErr = fmt.Errorf("something happened")
	)

	pet := &model.Pet{
		ID:          id,
		Name:        gofakeit.PetName(),
		Description: "description",
		Age:         int32(gofakeit.IntN(20)),
		OwnerID:     gofakeit.UUID(),
		CreatedAt:   gofakeit.Date(),
	}

	tests := []struct {
		name string
		want *model.Pet
		on   func()
	}{
		{
			name: "success_get_from_cache",
			want: pet,
			on: func() {
				s.petCache.GetByIDMock.
					Expect(s.ctx, id).
					Return(pet, nil)
			},
		},
		{
			name: "success_get_from_repository",
			want: pet,
			on: func() {
				s.petCache.GetByIDMock.
					Expect(s.ctx, id).
					Return(nil, petCacheErr)

				s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(pet, nil)
				s.petCache.SetMock.Expect(s.ctx, pet, s.ttl).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()

			tt.on()

			got, err := s.service.Get(s.ctx, id)
			s.Require().NotNil(pet)
			s.Require().EqualValues(pet, got)
			s.Require().NoError(err)
		})
	}
}

func (s *ServiceSuite) TestGet_Error() {
	var (
		id            = gofakeit.UUID()
		petCacheErr   = fmt.Errorf("something happened")
		repositoryErr = gofakeit.ErrorDatabase()
	)

	s.petCache.GetByIDMock.Expect(s.ctx, id).Return(nil, petCacheErr)
	s.petRepo.GetByIDMock.Expect(s.ctx, id).Return(nil, repositoryErr)

	pet, err := s.service.Get(s.ctx, id)
	s.Require().Nil(pet)
	s.Require().Error(err)
	s.Require().ErrorIs(err, repositoryErr)
}
