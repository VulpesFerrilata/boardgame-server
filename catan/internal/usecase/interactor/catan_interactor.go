package interactor

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/service"
	"github.com/VulpesFerrilata/boardgame-server/catan/internal/usecase/form"
)

type CatanInteractor interface {
	NewGame(ctx context.Context, userForm form.UserForm) error
	JoinGame(ctx context.Context, userForm form.UserForm) error
	StartGame(ctx context.Context, userForm form.UserForm) error
	LeaveGame(ctx context.Context, userForm form.UserForm) error
}

func NewCatanInteractor() CatanInteractor {
	return &catanInteractor{}
}

type catanInteractor struct {
	gameService service.GameService
}

func (ci catanInteractor) NewGame(ctx context.Context, userForm form.UserForm) error {
	game, err := ci.gameService.New(ctx)
	if err != nil {
		return err
	}

}
