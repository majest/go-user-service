package service

import (
	"github.com/majest/user-service/dao"
	"github.com/majest/user-service/server"
)

type UserService struct{}

func (s *UserService) FindOne(req *server.UserFindOneRequest) (*server.UserResponse, error) {
	udao := dao.NewUserDao()
	user, err := udao.FindOne(req.UserSearch)

	// if a pointer to a struct is passed, get the type of the dereferenced object

	return &server.UserResponse{
		User: user,
	}, err
}

func (s *UserService) Save(req *server.UserSaveRequest) (*server.UserSaveResponse, error) {
	udao := dao.NewUserDao()
	id, err := udao.Save(req.User)
	return &server.UserSaveResponse{&server.User{Id: id}}, err
}
