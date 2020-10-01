package service

import (
	"bytes"
	"context"
	"errors"

	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/repository"
	"gorm.io/gorm"

	server_errors "github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
)

func NewUserService(userRepo repository.UserRepository,
	translatorMiddleware *middleware.TranslatorMiddleware) *UserService {
	return &UserService{
		UserRepo:             userRepo,
		translatorMiddleware: translatorMiddleware,
	}
}

type UserService struct {
	UserRepo             repository.UserRepository
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (us UserService) ValidateLogin(ctx context.Context, user *model.User) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()
	userDB, err := us.UserRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return server_errors.NewNotFoundError("user")
		}
		return err
	}
	if !bytes.Equal(user.HashPassword, userDB.HashPassword) {
		fieldErr, _ := trans.T("validation-invalid", "password")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}

	return nil
}

func (us UserService) Validate(ctx context.Context, user *model.User) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()

	count, err := us.UserRepo.CountByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if count > 0 {
		fieldErr, _ := trans.T("validation-already-exists", "username")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}
