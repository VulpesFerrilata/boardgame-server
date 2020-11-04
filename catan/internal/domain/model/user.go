package model

import "github.com/VulpesFerrilata/grpc/protoc/user"

func NewUser(player *Player, userPb *user.UserResponse) *User {
	user := new(User)
	user.ID = uint(userPb.GetID())
	player.SetUser(user)
	return user
}

type User struct {
	ID       uint
	Username string
}
