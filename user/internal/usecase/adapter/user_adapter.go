package adapter

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/dto"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/form"

	"github.com/VulpesFerrilata/boardgame-server/library/pkg/validator"
)

type UserAdapter interface {
	ParseLogin(ctx context.Context, loginForm *form.LoginForm) (*model.User, error)
	ParseRegister(ctx context.Context, registerForm *form.RegisterForm) (*model.User, error)
	ParseUser(ctx context.Context, userForm *form.UserForm) (*model.User, error)
	ResponseUser(ctx context.Context, user *model.User) (*dto.UserDTO, error)
}

func NewUserAdapter(validate validator.Validate) UserAdapter {
	return &userAdapter{
		validate: validate,
	}
}

type userAdapter struct {
	validate validator.Validate
}

func (up userAdapter) ParseLogin(ctx context.Context, loginForm *form.LoginForm) (*model.User, error) {
	if err := up.validate.Struct(ctx, loginForm); err != nil {
		return nil, err
	}
	return loginForm.ToUser()
}

func (up userAdapter) ParseRegister(ctx context.Context, registerForm *form.RegisterForm) (*model.User, error) {
	if err := up.validate.Struct(ctx, registerForm); err != nil {
		return nil, err
	}
	return registerForm.ToUser()
}

func (up userAdapter) ParseUser(ctx context.Context, userForm *form.UserForm) (*model.User, error) {
	if err := up.validate.Struct(ctx, userForm); err != nil {
		return nil, err
	}
	return userForm.ToUser()
}

func (p userAdapter) ResponseUser(ctx context.Context, user *model.User) (*dto.UserDTO, error) {
	return dto.NewUserDTO(user)
}
