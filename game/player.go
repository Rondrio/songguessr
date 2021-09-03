package game

import "github.com/google/uuid"

type Player struct {
	id    string
	name  string
	lobby *Lobby
	conn  *Connection
}

func CreateNewPlayer(name string) *Player {
	return &Player{
		id:   uuid.NewString(),
		name: name,
	}
}

func (p *Player) CreateLobby(playerCount, rounds int, playlist, spotifyAuthToken string) *Lobby {
	lobby := Lobby{
		players: make([]*Player, 0),
		admin:   p,
		id:      uuid.NewString(),

		rounds:      rounds,
		playlist:    playlist,
		playerCount: playerCount,

		spotifyAuthToken: spotifyAuthToken,
	}

	lobbies = append(lobbies, &lobby)

	lobby.AddPlayer(p)

	return &lobby
}
