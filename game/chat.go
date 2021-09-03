package game

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Message struct {
	author  *Player
	content string
}

func (msg *Message) ParseMessage() error {
	var err error
	if msg.author.lobby != nil {
		log.Println(len(msg.author.lobby.players))
	}

	if !strings.HasPrefix(msg.content, "/") {
		if msg.author.lobby != nil {
			for _, p := range msg.author.lobby.players {
				log.Println(p.name)
				p.conn.conn.WriteMessage(1, []byte(fmt.Sprintf("%s: %s", msg.author.name, msg.content)))
			}
		}
	}

	parts := strings.Fields(msg.content)

	switch parts[0] {
	case "/join":
		lobby, err := FindLobby(parts[1])
		if err != nil {
			return err
		}

		if len(lobby.players) > lobby.playerCount {
			msg.author.conn.conn.WriteMessage(1, []byte("Lobby already full"))
			return nil
		}

		lobby.AddPlayer(msg.author)
		msg.author.lobby = lobby
		log.Println(msg.author.id, "joining lobby with id", lobby.id)
		log.Println(len(msg.author.lobby.players))

	case "/name":
		msg.author.name = strings.Join(parts[1:], " ")
		log.Println(msg.author.id, "changing name to", msg.author.name)

	case "/create":
		playerCount := 10
		if !(len(parts) < 2) {
			playerCount, err = strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
		}

		rounds := 5
		if !(len(parts) < 3) {
			rounds, err = strconv.Atoi(parts[2])
			if err != nil {
				return err
			}
		}

		playlist := "Top 50 - Germany"
		if !(len(parts) < 4) {
			playlist = parts[3]
		}

		lobby := msg.author.CreateLobby(playerCount, rounds, playlist, "")
		msg.author.lobby = lobby
		log.Println(msg.author.id, "creating lobby with id", lobby.id)
	}

	return nil
}
