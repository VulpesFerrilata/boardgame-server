package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"

	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
)

type UserRepository interface {
	CountByUsername(ctx context.Context, username string) (int, error)
	GetById(ctx context.Context, id uint) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	FindAll(context.Context) ([]*model.User, error)
	Insert(context.Context, *model.User) error
}

func NewUserRepository(dbContext *db.DbContext) UserRepository {
	return &userRepository{
		dbContext: dbContext,
	}
}

type userRepository struct {
	dbContext *db.DbContext
}

func (ur userRepository) CountByUsername(ctx context.Context, username string) (int, error) {
	var count int64
	user := new(model.User)
	return int(count), ur.dbContext.GetDB(ctx).Find(user, "username = ?", username).Count(&count).Error
}

func (ur userRepository) GetById(ctx context.Context, id uint) (*model.User, error) {
	user := new(model.User)
	return user, ur.dbContext.GetDB(ctx).First(user, id).Error
}

func (ur userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	return user, ur.dbContext.GetDB(ctx).First(user, "username = ?", username).Error
}

func (ur userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)
	return users, ur.dbContext.GetDB(ctx).Find(&users).Error
}

func (ur userRepository) Insert(ctx context.Context, user *model.User) error {
	return ur.dbContext.GetDB(ctx).Create(user).Error
}
