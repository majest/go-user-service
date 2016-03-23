package server

type UserService interface {
	FindOne(req *UserFindOneRequest) (*UserResponse, error)
	//FindAll(userSearch *UserRequest, limit int, lastId string) *UsersResponse
}

type User struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Address   string `json:"address,omitempty"`
	Street    string `json:"street,omitempty"`
	PostCode  string `json:"postCode,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type UserFindOneRequest struct {
	UserSearch *UserSearch
}

type UserResponse struct {
	User *User `json:"user,omitempty"`
}

type UserSaveRequest struct {
	User *User
}

type UserSaveResponse struct {
	User *User
}

type UsersResponse struct {
	Users   []*User `json:"users,omitempty"`
	LastKey string  `json:"lastKey,omitempty"`
}

type UserSearch struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	PostCode  string
}
