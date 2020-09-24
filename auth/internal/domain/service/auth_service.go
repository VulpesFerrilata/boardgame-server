package service

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/repository"
	server_errors "github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"

	"gorm.io/gorm"
)

func NewAuthService(authRepo repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepo: authRepo,
	}
}

type AuthService struct {
	AuthRepo repository.AuthRepository
}

func (as AuthService) Validate(ctx context.Context, token *model.Token) error {
	validationErrs := new(server_errors.ValidationError)
	count, err := as.AuthRepo.CountByJti(ctx, token.Jti)
	if err != nil {
		return err
	}
	if count > 0 {
		validationErrs.Append("jti is already exists")
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}

func (as AuthService) ValidateAuthenticate(ctx context.Context, token *model.Token) error {
	validationErrs := new(server_errors.ValidationError)
	tokenDB, err := as.AuthRepo.GetById(ctx, token.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			validationErrs.Append("id is invalid")
			return validationErrs
		}
		return err
	}
	if token.Jti != tokenDB.Jti {
		validationErrs.Append("jti is invalid")
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}
