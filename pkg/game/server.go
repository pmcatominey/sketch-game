package game

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	nameRegex = regexp.MustCompile(`^\w+$`)

	ErrNameRegex          = errors.New("player name must match regex ^\\w+$")
	ErrPlayerNameConflict = errors.New("player already connected with provided name")
)

type Server struct {
	games    map[string]*game
	gamesMu  sync.RWMutex
	upgrader *websocket.Upgrader
}

func NewServer(openCors bool) *Server {
	s := &Server{
		games:    map[string]*game{},
		gamesMu:  sync.RWMutex{},
		upgrader: &websocket.Upgrader{},
	}

	if openCors {
		s.upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade WS connection: %s\n", err)
		return
	}

	var (
		key        = r.URL.Query().Get("key")
		playerName = r.URL.Query().Get("name")
	)

	if !nameRegex.MatchString(playerName) {
		conn.WriteJSON(replyErr(ErrNameRegex))
		conn.Close()
		return
	}

	s.gamesMu.Lock()
	defer s.gamesMu.Unlock()

	game, ok := s.games[key]
	if !ok {
		game = newGame()
		s.games[game.key] = game
	}

	if _, ok = game.sessions[playerName]; ok {
		conn.WriteJSON(replyErr(ErrPlayerNameConflict))
		conn.Close()
		return
	}

	go game.handle(playerName, conn)
}
