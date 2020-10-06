package service

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	server_errors "github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
)

type UserService interface {
	GetUserRepository() repository.UserRepository
	ValidateLogin(ctx context.Context, user *model.User, plainPassword string) error
	Validate(ctx context.Context, user *model.User) error
}

func NewUserService(userRepository repository.UserRepository,
	translatorMiddleware *middleware.TranslatorMiddleware) UserService {
	return &userService{
		userRepository:       userRepository,
		translatorMiddleware: translatorMiddleware,
	}
}

type userService struct {
	userRepository       repository.UserRepository
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (us userService) GetUserRepository() repository.UserRepository {
	return us.userRepository
}

func (us userService) ValidateLogin(ctx context.Context, user *model.User, plainPassword string) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()
	userDB, err := us.userRepository.GetByUsername(ctx, user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return server_errors.NewNotFoundError("user")
		}
		return err
	}
	if err := bcrypt.CompareHashAndPassword(userDB.HashPassword, []byte(plainPassword)); err != nil {
		fieldErr, _ := trans.T("validation-invalid", "password")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}

	return nil
}

func (us userService) Validate(ctx context.Context, user *model.User) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()

	count, err := us.userRepository.CountByUsername(ctx, user.Username)
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
