package converter

import (
	"github.com/escoutdoor/kotopes/auth/internal/model"
	pb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbFromUser(in *model.User) *pb.User {
	return &pb.User{
		Id:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
		City:      in.City,
		Country:   in.Country,
		CreatedAt: timestamppb.New(in.CreatedAt),
	}
}

func ToPbFromUsers(in []*model.User) []*pb.User {
	var users []*pb.User
	for _, u := range in {
		users = append(users, ToPbFromUser(u))
	}

	return users
}

func ToUpdateUserFromPb(pb *pb.UpdateRequest) *model.UpdateUser {
	out := &model.UpdateUser{
		ID: pb.GetId(),
	}

	if pb.GetFirstName() != nil {
		firstName := pb.FirstName.GetValue()
		out.FirstName = &firstName
	}
	if pb.GetLastName() != nil {
		lastName := pb.LastName.GetValue()
		out.LastName = &lastName
	}
	if pb.GetEmail() != nil {
		email := pb.Email.GetValue()
		out.Email = &email
	}
	if pb.GetPassword() != nil {
		password := pb.Password.GetValue()
		out.Password = &password
	}
	if pb.GetPhone() != nil {
		phone := pb.Phone.GetValue()
		out.Phone = &phone
	}
	if pb.GetCity() != nil {
		city := pb.City.GetValue()
		out.City = &city
	}
	if pb.GetCountry() != nil {
		country := pb.Country.GetValue()
		out.Country = &country
	}

	return out
}
