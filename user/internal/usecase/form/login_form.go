package form

import (
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"

	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Username string `name:"username"`
	Password string `name:"password"`
}

func (lf LoginForm) ToUser() (*model.User, error) {
	user := new(model.User)
	user.Username = lf.Username
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(lf.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.HashPassword = hashPassword
	return user, nil
}
