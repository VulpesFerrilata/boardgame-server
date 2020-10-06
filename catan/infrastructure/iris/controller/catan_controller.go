package controller

import "github.com/kataras/iris/v12/websocket"

type CatanController interface {
}

func NewCatanController() CatanController {
	return &catanController{}
}

type catanController struct {
	*websocket.NSConn `stateless:"true"`
}

func (cc catanController) Namespace() string {
	return "default"
}
