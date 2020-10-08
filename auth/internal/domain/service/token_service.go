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

type AuthService interface {
	GetAuthRepository() repository.ReadOnlyTokenRepository
	ValidateAuthenticate(ctx context.Context, token *model.Token) error
	CreateOrUpdate(ctx context.Context, token *model.Token) error
}

func NewAuthService(tokenRepository repository.TokenRepository,
	translatorMiddleware *middleware.TranslatorMiddleware) AuthService {
	return &authService{
		tokenRepository:      tokenRepository,
		translatorMiddleware: translatorMiddleware,
	}
}

type authService struct {
	tokenRepository      repository.TokenRepository
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (as authService) GetAuthRepository() repository.ReadOnlyTokenRepository {
	return as.tokenRepository
}

func (as authService) ValidateAuthenticate(ctx context.Context, token *model.Token) error {
	trans := as.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()
	tokenDB, err := as.tokenRepository.GetByUserId(ctx, token.UserID)
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

func (as authService) CreateOrUpdate(ctx context.Context, token *model.Token) error {
	tokenDB, err := as.tokenRepository.GetByUserId(ctx, token.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	token.ID = tokenDB.ID
	return as.tokenRepository.Save(ctx, token)
}
