package dto

type PlayerDTO struct {
	ID               int
	User             *UserDTO
	Color            string
	TurnOrder        int
	IsConfirmed      bool
	ResourceCards    []*ResourceCardDTO
	DevelopmentCards []*DevelopmentCardDTO
}
