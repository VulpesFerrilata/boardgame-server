package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/form"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
)

type TokenRequest struct {
	*auth.TokenRequest
}

func (tr TokenRequest) ToTokenForm() *form.TokenForm {
	tokenForm := new(form.TokenForm)
	tokenForm.Token = tr.GetToken()
	return tokenForm
}
