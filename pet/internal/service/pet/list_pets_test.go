package pet

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (s *ServiceSuite) TestListPets_Success() {
	var (
		limit  = int32(gofakeit.IntRange(1, 50))
		offset = int32(gofakeit.IntRange(1, 50))

		petIDs = make([]string, limit)
		pets   = make([]*model.Pet, limit)
	)

	for i := 0; i < int(limit); i++ {
		id := gofakeit.UUID()
		petIDs[i] = id

		pets[i] = &model.Pet{
			ID:          id,
			Name:        gofakeit.PetName(),
			Description: gofakeit.Paragraph(1, 1, 1, ""),
			Age:         int32(gofakeit.IntRange(1, 20)),
			OwnerID:     gofakeit.UUID(),
			CreatedAt:   gofakeit.Date(),
		}
	}

	in := &model.ListPets{
		Limit:  limit,
		Offset: offset,
		PetIDs: petIDs,
	}

	s.petRepo.ListPetsMock.Expect(s.ctx, in).Return(pets, nil)

	got, err := s.service.ListPets(s.ctx, in)
	s.Require().NotNil(got)
	s.Require().EqualValues(pets, got)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestListPets_Error() {
	var (
		limit  = int32(gofakeit.IntRange(1, 50))
		offset = int32(gofakeit.IntRange(1, 50))

		petIDs = make([]string, limit)
		pets   = make([]*model.Pet, limit)

		petRepoError = fmt.Errorf("pet repository error")
	)

	for i := 0; i < int(limit); i++ {
		id := gofakeit.UUID()
		petIDs[i] = id

		pets[i] = &model.Pet{
			ID:          id,
			Name:        gofakeit.PetName(),
			Description: gofakeit.Paragraph(1, 1, 1, ""),
			Age:         int32(gofakeit.IntRange(1, 20)),
			OwnerID:     gofakeit.UUID(),
			CreatedAt:   gofakeit.Date(),
		}
	}

	in := &model.ListPets{
		Limit:  limit,
		Offset: offset,
		PetIDs: petIDs,
	}

	s.petRepo.ListPetsMock.Expect(s.ctx, in).Return(nil, petRepoError)

	got, err := s.service.ListPets(s.ctx, in)
	s.Require().Nil(got)
	s.Require().Error(err)
	s.Require().ErrorIs(err, petRepoError)
}
