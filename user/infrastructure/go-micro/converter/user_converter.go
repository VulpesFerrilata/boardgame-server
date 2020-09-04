package converter

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/dto"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"
)

func ConvertUserRequestPbToUserForm(userRequestPb *user.UserRequest) *form.UserForm {
	userForm := new(form.UserForm)
	userForm.ID = int(userRequestPb.ID)
	return userForm
}

func ConvertCredentialRequestPbToLoginForm(credentialRequestPb *user.CredentialRequest) *form.LoginForm {
	loginForm := new(form.LoginForm)
	loginForm.Username = credentialRequestPb.Username
	loginForm.Password = credentialRequestPb.Password
	return loginForm
}

func ConvertUserDtoToUserResponsePb(userDTO *dto.UserDTO) *user.UserResponse {
	userResponsePb := new(user.UserResponse)
	userResponsePb.ID = int64(userDTO.ID)
	return userResponsePb
}
