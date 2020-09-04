package request

import "github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"

type RegisterRequest struct {
	LoginRequest
	RepeatPassword string `json:"repeatPassword"`
}

func (rr RegisterRequest) ToInteractorRegisterForm() *form.RegisterForm {
	registerForm := new(form.RegisterForm)
	registerForm.LoginForm = rr.LoginRequest.ToInteractorLoginForm()
	registerForm.RepeatPassword = rr.RepeatPassword
	return registerForm
}
