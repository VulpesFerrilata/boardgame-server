package request

type GameRequest struct {
	ID int `json:"id" validate:"required"`
}
