package game

import (
	"errors"
)

var lobbies []*Lobby

func FindLobby(id string) (*Lobby, error) {
	for _, lobby := range lobbies {
		if lobby.id == id {
			return lobby, nil
		}
	}

	return nil, errors.New("lobby not found")
}

func startGame(lobby *Lobby) {

}
