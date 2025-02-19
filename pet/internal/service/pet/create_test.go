package pet

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (s *ServiceSuite) TestCreate_Success() {
	var (
		id          = gofakeit.UUID()
		name        = gofakeit.PetName()
		description = gofakeit.Paragraph(5, 5, 5, "")
		age         = int32(gofakeit.IntN(20))
	)

	in := &model.CreatePet{
		Name:        name,
		Description: description,
		Age:         age,
		OwnerID:     gofakeit.UUID(),
	}

	s.petRepo.CreateMock.Expect(s.ctx, in).Return(id, nil)

	createPetID, err := s.service.Create(s.ctx, in)
	s.Require().NotEmpty(createPetID)
	s.Require().Equal(id, createPetID)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestCreate_Error() {
	var (
		name        = gofakeit.PetName()
		description = gofakeit.Paragraph(5, 5, 5, "")
		age         = int32(gofakeit.IntN(20))

		repoErr = gofakeit.Error()
	)

	in := &model.CreatePet{
		Name:        name,
		Description: description,
		Age:         age,
		OwnerID:     gofakeit.UUID(),
	}

	s.petRepo.CreateMock.Expect(s.ctx, in).Return("", repoErr)

	id, err := s.service.Create(s.ctx, in)
	s.Require().Empty(id)
	s.Require().Error(err)
	s.Require().ErrorIs(err, repoErr)
}

// func (s *ServiceSuite) TestCreateValidation_Success() {
// 	in := &model.CreatePet{
// 		Name:        gofakeit.PetName(),
// 		Description: gofakeit.Paragraph(5, 5, 5, ""),
// 		Age:         int32(gofakeit.IntN(20)),
// 		OwnerID:     gofakeit.UUID(),
// 	}
//
// 	err := validateCreateRequest(in)
// 	s.Require().Nil(err)
// 	s.Require().NoError(err)
// }
//
// func (s *ServiceSuite) TestCreateValidation_Error() {
// 	tests := []struct {
// 		name        string
// 		in          *model.CreatePet
// 		expectedErr error
// 	}{
// 		{
// 			name:        "bad_empty_name",
// 			in:          &model.CreatePet{},
// 			expectedErr: model.ErrEmptyPetName,
// 		},
// 		{
// 			name:        "bad_empty_description",
// 			in:          &model.CreatePet{Name: gofakeit.PetName()},
// 			expectedErr: model.ErrEmptyPetDescription,
// 		},
// 		{
// 			name: "bad_invalid_age",
// 			in: &model.CreatePet{
// 				Name:        gofakeit.PetName(),
// 				Description: gofakeit.Paragraph(5, 5, 5, ""),
// 				Age:         -1,
// 			},
// 			expectedErr: model.ErrInvalidPetAge,
// 		},
// 		{
// 			name: "bad_excessive_age",
// 			in: &model.CreatePet{
// 				Name:        gofakeit.PetName(),
// 				Description: gofakeit.Paragraph(5, 5, 5, ""),
// 				Age:         10000,
// 			},
// 			expectedErr: model.ErrExcessivePetAge,
// 		},
// 		{
// 			name: "bad_empty_owner_id",
// 			in: &model.CreatePet{
// 				Name:        gofakeit.PetName(),
// 				Description: gofakeit.Paragraph(5, 5, 5, ""),
// 				Age:         int32(gofakeit.IntN(20)),
// 			},
// 			expectedErr: model.ErrEmptyOwnerID,
// 		},
// 		{
// 			name: "bad_invalid_owner_id",
// 			in: &model.CreatePet{
// 				Name:        gofakeit.PetName(),
// 				Description: gofakeit.Paragraph(5, 5, 5, ""),
// 				Age:         int32(gofakeit.IntN(20)),
// 				OwnerID:     gofakeit.MovieName(),
// 			},
// 			expectedErr: model.ErrInvalidOwnerID,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		s.T().Run(tt.name, func(t *testing.T) {
// 			err := validateCreateRequest(tt.in)
// 			s.Require().Error(err)
// 			s.Require().EqualError(err, tt.expectedErr.Error())
// 		})
// 	}
// }
