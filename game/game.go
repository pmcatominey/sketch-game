package game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pmcatominey/sketch-game/state"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type game struct {
	key string

	mutex        *sync.Mutex
	endTimerOnce *sync.Once

	state *state.State

	sessions map[string]*websocket.Conn
}

type gameSession struct {
	conn *websocket.Conn
}

func newGame() *game {
	return &game{
		mutex:        &sync.Mutex{},
		endTimerOnce: &sync.Once{},

		key:   genKey(16),
		state: state.New(),

		sessions: map[string]*websocket.Conn{},
	}
}

func (g *game) handle(playerName string, conn *websocket.Conn) {
	g.mutex.Lock()

	g.sessions[playerName] = conn
	g.state.Connect(playerName)

	conn.WriteJSON(replyHello(g.key))

	g.broadcastState()
	g.broadcast(replyEvent(fmt.Sprintf("%s joined", playerName)))
	g.mutex.Unlock()

	for {
		var cmd command
		if err := conn.ReadJSON(&cmd); err != nil {
			break
		}

		updated := false
		updateMsg := ""

		g.mutex.Lock()
		switch cmd.Type {
		case "chat":
			g.broadcast(replyChat(playerName, cmd.Data["message"]))
		case "submitImage":
			updated = g.state.SubmitImage(playerName, cmd.Data["image"])
			updateMsg = fmt.Sprintf("%s has finished drawing", playerName)
		case "submitGuess":
			updated = g.state.SubmitGuess(playerName, cmd.Data["guess"])
		case "markGuessCorrect":
			updated = g.state.MarkGuess(playerName, cmd.Data["player"], true)
			updateMsg = fmt.Sprintf("%s has guessed correctly!", playerName)
		case "markGuessIncorrect":
			updated = g.state.MarkGuess(playerName, cmd.Data["player"], false)
		}

		if updated {
			g.broadcastState()

			if updateMsg != "" {
				g.broadcast(replyEvent(updateMsg))
			}

			if g.state.Phase == "result" {
				go g.endTimerOnce.Do(func() {
					time.Sleep(time.Second * 10)

					g.mutex.Lock()
					defer g.mutex.Unlock()

					if g.state.EndRound() {
						g.broadcastState()
					}

					g.endTimerOnce = &sync.Once{}
				})
			}
		}
		g.mutex.Unlock()
	}

	g.mutex.Lock()
	conn.Close()

	if g.state.Disconnect(playerName) {
		delete(g.sessions, playerName)
		g.broadcastState()
		g.broadcast(replyEvent(fmt.Sprintf("%s has left", playerName)))
	}

	g.mutex.Unlock()
}

func (g *game) broadcast(r reply) {
	for _, session := range g.sessions {
		session.WriteJSON(r)
	}
}

func (g *game) broadcastState() {
	g.broadcast(replyState(g.state))
}

func genKey(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[random.Intn(len(letterRunes))]
	}
	return string(b)
}
