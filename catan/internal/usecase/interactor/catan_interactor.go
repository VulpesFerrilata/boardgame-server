package interactor

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/usecase/form"
)

type CatanInteractor interface {
	NewGame(ctx context.Context, userForm form.UserForm)
	JoinGame(ctx context.Context, userForm form.UserForm)
}

func NewCatanInteractor() CatanInteractor {
	return &catanInteractor{}
}

type catanInteractor struct {
}

func (ci catanInteractor) NewGame(ctx context.Context, userForm form.UserForm) {

}
