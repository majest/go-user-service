package server

import (
	"fmt"

	"github.com/majest/user-service/pb"
	"golang.org/x/net/context"
)

func DecodeFindOneRequest(ctx context.Context, req interface{}) (interface{}, error) {
	fmt.Println(req)
	domainRequest := req.(*pb.UserFindOneRequest)

	return &UserFindOneRequest{&UserSearch{
		Id:        domainRequest.UserSearch.Id,
		FirstName: domainRequest.UserSearch.FirstName,
		LastName:  domainRequest.UserSearch.LastName,
		Email:     domainRequest.UserSearch.Email,
		PostCode:  domainRequest.UserSearch.PostCode,
	}}, nil
}

func EncodeFindOneResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	fmt.Println(resp)
	domainResponse := resp.(*UserResponse)

	if domainResponse.User != nil {
		return &pb.UserResponse{
			User: &pb.User{
				Id:        domainResponse.User.Id,
				FirstName: domainResponse.User.FirstName,
				LastName:  domainResponse.User.LastName,
				Email:     domainResponse.User.Email,
				Address:   domainResponse.User.Address,
				Street:    domainResponse.User.Street,
				PostCode:  domainResponse.User.PostCode,
				City:      domainResponse.User.City,
				Country:   domainResponse.User.Country,
				Phone:     domainResponse.User.Phone,
			},
		}, nil
	}

	return &pb.UserResponse{}, nil
}
