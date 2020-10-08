package form

type PlayerForm struct {
	ID          int
	User        UserForm
	Color       string
	TurnOrder   int
	IsConfirmed bool
}
