package transaction

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

type transactionKey struct {}

func NewTransactionMiddleware(db *gorm.DB) *TransactionMiddleware {
	return &TransactionMiddleware{
		db: db,
	}
}

type TransactionMiddleware struct {
	db *gorm.DB
}

func (tm TransactionMiddleware) Serve(ctx iris.Context) {
	r := ctx.Request()
	requestCtx := r.Context()
	tx := tm.db.BeginTx(requestCtx, &sql.TxOptions{})

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	requestCtx = context.WithValue(requestCtx, transactionKey{}, tx)
	r.WithContext(requestCtx)
	ctx.ResetRequest(r)

	ctx.Next()
}

func (tm TransactionMiddleware) Get(ctx context.Context) (*gorm.DB, bool) {
	tx, found := ctx.Value(transactionKey{}).(*gorm.DB)
	return tx, found
}
