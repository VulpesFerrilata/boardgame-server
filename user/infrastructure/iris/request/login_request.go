package request

import "github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr LoginRequest) ToInteractorLoginForm() *form.LoginForm {
	loginForm := new(form.LoginForm)
	loginForm.Username = lr.Username
	loginForm.Password = lr.Password
	return loginForm
}
