package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"
)

type CredentialRequest struct {
	*user.CredentialRequest
}

func (cr CredentialRequest) ToLoginForm() *form.LoginForm {
	loginForm := new(form.LoginForm)
	loginForm.Username = cr.Username
	loginForm.Password = cr.Password
	return loginForm
}
