package game

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{} // use default options

func InitServer() {
	c := Connection{}

	http.HandleFunc("/connect", c.handleConnect)

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func (c *Connection) handleConnect(w http.ResponseWriter, r *http.Request) {
	var err error
	c.conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	log.Println("client connected")

	player := CreateNewPlayer("")
	player.conn = c

	defer c.conn.Close()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		msg := &Message{
			author:  player,
			content: string(message),
		}

		err = msg.ParseMessage()
		if err != nil {
			log.Println(err)
		}
	}
}

/*
func initServer() {
	router := mux.NewRouter()

	router.HandleFunc("/createLobby", handleCreateLobby).Methods("POST")
	router.HandleFunc("/joinLobby", handleJoinLobby).Methods("POST")

	log.Fatal(http.ListenAndServe(":420", router))
}

type CreateLobbyBody struct {
	PlayerName  string
	PlayerCount int
	Rounds      int
	Playlist    string

	SpotifyAuthToken string
}

func handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var body CreateLobbyBody
	r.Body.Read(b)

	err := json.Unmarshal(b, &body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	player := game.CreateNewPlayer(body.PlayerName)

	lobby := player.CreateLobby(body.PlayerCount, body.Rounds, body.Playlist, body.SpotifyAuthToken)
	w.Write([]byte(lobby.GetId()))
	w.WriteHeader(200)
}

type JoinLobbyBody struct {
	PlayerName string
	LobbyId    string
}

func handleJoinLobby(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var body JoinLobbyBody
	r.Body.Read(b)

	err := json.Unmarshal(b, &body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	player := game.CreateNewPlayer(body.PlayerName)

	lobby, err := game.FindLobby(body.LobbyId)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	lobby.AddPlayer(player)
}
*/
