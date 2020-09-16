package db

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"

	"github.com/jinzhu/gorm"
)

func NewDbContext(db *gorm.DB, transactionMiddleware *middleware.TransactionMiddleware) *DbContext {
	return &DbContext{
		db:                    db,
		transactionMiddleware: transactionMiddleware,
	}
}

type DbContext struct {
	db                    *gorm.DB
	transactionMiddleware *middleware.TransactionMiddleware
}

func (dc *DbContext) GetDB(ctx context.Context) *gorm.DB {
	tx, found := dc.transactionMiddleware.Get(ctx)
	if found {
		return tx
	}
	return dc.db
}
