package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"
)

type UserRequest struct {
	*user.UserRequest
}

func (ur UserRequest) ToUserForm() *form.UserForm {
	userForm := new(form.UserForm)
	userForm.ID = int(ur.ID)
	return userForm
}
