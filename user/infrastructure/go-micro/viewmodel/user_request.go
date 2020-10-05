package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"
)

func NewUserRequest(userRequestPb *user.UserRequest) *UserRequest {
	return &UserRequest{
		userRequestPb: userRequestPb,
	}
}

type UserRequest struct {
	userRequestPb *user.UserRequest
}

func (ur UserRequest) ToUserForm() *form.UserForm {
	userForm := new(form.UserForm)
	userForm.ID = int(ur.userRequestPb.GetID())
	return userForm
}
