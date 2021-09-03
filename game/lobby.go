package game

type Lobby struct {
	id      string
	players []*Player
	admin   *Player

	rounds      int
	playlist    string
	playerCount int

	spotifyAuthToken string
}

func (l *Lobby) AddPlayer(p *Player) {
	l.players = append(l.players, p)
}

/*
func (l *Lobby) RemovePlayer(p *Player) {
	for i := 0; i < len(l.players); i++ {
		if l.players[i] == p {
			l.players = append(l.players[:i], l.players[i+1:]...)
		}
	}
}
*/

func (l *Lobby) startGame() {

}
