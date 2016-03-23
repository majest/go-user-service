package server

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakeFindOneEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println(request)
		req := request.(*UserFindOneRequest)
		user, err := svc.FindOne(req)
		return user, err
	}
}
