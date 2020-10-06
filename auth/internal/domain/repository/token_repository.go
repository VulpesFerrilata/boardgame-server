package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
)

type TokenRepository interface {
	CountByJti(ctx context.Context, jti string) (int, error)
	GetByUserId(ctx context.Context, userId uint) (*model.Token, error)
	GetByJti(ctx context.Context, jti string) (*model.Token, error)
	Save(context.Context, *model.Token) error
}

func NewTokenRepository(dbContext *db.DbContext) TokenRepository {
	return &tokenRepository{
		dbContext: dbContext,
	}
}

type tokenRepository struct {
	dbContext *db.DbContext
}

func (tr tokenRepository) CountByJti(ctx context.Context, jti string) (int, error) {
	var count int64
	token := new(model.Token)
	return int(count), tr.dbContext.GetDB(ctx).Find(token, "jti = ?", jti).Count(&count).Error
}

func (tr tokenRepository) GetByUserId(ctx context.Context, userId uint) (*model.Token, error) {
	token := new(model.Token)
	return token, tr.dbContext.GetDB(ctx).First(token, userId).Error
}

func (tr tokenRepository) GetByJti(ctx context.Context, jti string) (*model.Token, error) {
	token := new(model.Token)
	return token, tr.dbContext.GetDB(ctx).First(token, "jti = ?", jti).Error
}

func (tr tokenRepository) Save(ctx context.Context, token *model.Token) error {
	return tr.dbContext.GetDB(ctx).Save(token).Error
}
