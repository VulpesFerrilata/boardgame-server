package form

import (
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
)

type UserForm struct {
	ID int `name:"id"`
}

func (uf UserForm) ToUser() (*model.User, error) {
	user := new(model.User)
	user.ID = uint(uf.ID)
	return user, nil
}
