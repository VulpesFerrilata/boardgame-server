package service

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/repository"
	server_errors "github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"gorm.io/gorm"
)

func NewAuthService(authRepo repository.AuthRepository,
	translatorMiddleware *middleware.TranslatorMiddleware) *AuthService {
	return &AuthService{
		AuthRepo:             authRepo,
		translatorMiddleware: translatorMiddleware,
	}
}

type AuthService struct {
	AuthRepo             repository.AuthRepository
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (as AuthService) ValidateAuthenticate(ctx context.Context, token *model.Token) error {
	trans := as.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()
	tokenDB, err := as.AuthRepo.GetByUserId(ctx, token.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return server_errors.NewNotFoundError("token")
		}
		return err
	}
	if token.Jti != tokenDB.Jti {
		fieldErr, _ := trans.T("validation-invalid", "jti")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}

func (as AuthService) CreateOrUpdate(ctx context.Context, token *model.Token) error {
	tokenDB, err := as.AuthRepo.GetByUserId(ctx, token.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	token.ID = tokenDB.ID
	return as.AuthRepo.CreateOrUpdate(ctx, token)
}
