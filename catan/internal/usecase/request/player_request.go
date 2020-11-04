package request

type PlayerRequest struct {
	ID          int
	GameID      int
	UserID      int
	Color       string
	TurnOrder   int
	IsConfirmed bool
}
