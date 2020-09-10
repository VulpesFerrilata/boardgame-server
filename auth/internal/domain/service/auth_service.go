package service

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/repository"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"

	"github.com/jinzhu/gorm"
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
	validationErrs := new(errors.ValidationError)
	tokenDB, err := as.AuthRepo.GetByJti(ctx, token.Jti)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if tokenDB != nil {
		validationErrs.Append("jti is already exists")
	}
	return validationErrs
}

func (as AuthService) ValidateAuthenticate(ctx context.Context, token *model.Token) error {
	validationErrs := new(errors.ValidationError)
	tokenDB, err := as.AuthRepo.GetById(ctx, token.ID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			validationErrs.Append("id is invalid")
			return validationErrs
		}
		return err
	}
	if token.Jti != tokenDB.Jti {
		validationErrs.Append("jti is invalid")
	}
	return validationErrs
}
