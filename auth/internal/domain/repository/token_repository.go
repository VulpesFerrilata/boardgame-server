package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
)

type AuthRepository interface {
	CountByJti(ctx context.Context, jti string) (int, error)
	GetById(ctx context.Context, id uint) (*model.Token, error)
	GetByJti(ctx context.Context, jti string) (*model.Token, error)
	CreateOrUpdate(context.Context, *model.Token) error
}

func NewAuthRepository(dbContext *db.DbContext) AuthRepository {
	return &authRepository{
		dbContext: dbContext,
	}
}

type authRepository struct {
	dbContext *db.DbContext
}

func (ar authRepository) CountByJti(ctx context.Context, jti string) (int, error) {
	count := 0
	token := new(model.Token)
	return count, ar.dbContext.GetDB(ctx).Find(token, "jti = ?", jti).Count(&count).Error
}

func (ar authRepository) GetById(ctx context.Context, id uint) (*model.Token, error) {
	token := new(model.Token)
	return token, ar.dbContext.GetDB(ctx).First(token, id).Error
}

func (ar authRepository) GetByJti(ctx context.Context, jti string) (*model.Token, error) {
	token := new(model.Token)
	return token, ar.dbContext.GetDB(ctx).First(token, "jti = ?", jti).Error
}

func (ar authRepository) CreateOrUpdate(ctx context.Context, token *model.Token) error {
	return ar.dbContext.GetDB(ctx).Save(token).Error
}
