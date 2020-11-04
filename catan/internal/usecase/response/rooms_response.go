package response

import "github.com/VulpesFerrilata/catan/internal/domain/model"

func NewRoomsResponse(rooms ...*model.Room) RoomsResponse {
	var roomsResponse RoomsResponse

	for _, room := range rooms {
		roomResponse := NewRoomResponse(room)
		roomsResponse = append(roomsResponse, roomResponse)
	}

	return roomsResponse
}

type RoomsResponse []*RoomResponse
