package service

import (
	"bytes"
	"context"

	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/repository"

	server_errors "github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"

	"errors"

	"gorm.io/gorm"
)

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type UserService struct {
	UserRepo repository.UserRepository
}

func (us UserService) ValidateLogin(ctx context.Context, user *model.User) error {
	validationErrs := new(server_errors.ValidationError)

	userDB, err := us.UserRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			validationErrs.Append("username is invalid")
			return validationErrs
		}
		return err
	}
	if !bytes.Equal(user.HashPassword, userDB.HashPassword) {
		validationErrs.Append("password is invalid")
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}

	return nil
}

func (us UserService) Validate(ctx context.Context, user *model.User) error {
	validationErrs := new(server_errors.ValidationError)

	count, err := us.UserRepo.CountByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if count > 0 {
		validationErrs.Append("username is already exists")
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}
