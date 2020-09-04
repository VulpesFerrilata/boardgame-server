package form

import "github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"

type RegisterForm struct {
	*LoginForm
	RepeatPassword string `name:"repeat password"`
}

func (rf RegisterForm) ToUser() (*model.User, error) {
	user, err := rf.LoginForm.ToUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}
