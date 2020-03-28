package state

import (
	"sort"
)

const (
	phaseWaiting  = "waiting"
	phaseDrawing  = "drawing"
	phaseGuessing = "guessing"
	phaseResult   = "result"
)

type State struct {
	Phase   string                  `json:"phase"`
	Players map[string]*PlayerState `json:"players"`

	DrawingPlayer string `json:"drawingPlayer"`
	Image         string `json:"image"`

	CorrectPlayer string `json:"correctPlayer"`
	CorrectAnswer string `json:"correctAnswer"`
}

type PlayerState struct {
	Points         int    `json:"points"`
	Guess          string `json:"guess"`
	GuessIncorrect bool   `json:"guessIncorrect"`
}

func New() *State {
	return &State{
		Phase:   phaseWaiting,
		Players: map[string]*PlayerState{},
	}
}

func (s *State) Connect(playerName string) bool {
	_, ok := s.Players[playerName]
	if !ok {
		s.Players[playerName] = &PlayerState{}

		if s.Phase == phaseWaiting && len(s.Players) > 1 {
			s.nextDrawer()
		}
	}

	return !ok
}

func (s *State) Disconnect(playerName string) bool {
	_, ok := s.Players[playerName]
	if ok {
		delete(s.Players, playerName)

		if (s.Phase == phaseDrawing || s.Phase == phaseGuessing) && s.DrawingPlayer == playerName {
			s.nextDrawer()
		}
	}

	return ok
}

func (s *State) SubmitImage(playerName, image string) bool {
	if s.Phase != phaseDrawing || image == "" {
		return false
	}

	_, ok := s.Players[playerName]
	if !ok || s.DrawingPlayer != playerName {
		return false
	}

	s.Image = image
	s.Phase = phaseGuessing
	return true
}

func (s *State) SubmitGuess(playerName, guess string) bool {
	if s.Phase != phaseGuessing || guess == "" {
		return false
	}

	player, ok := s.Players[playerName]
	if !ok || s.DrawingPlayer == playerName {
		return false
	}

	player.Guess = guess
	player.GuessIncorrect = false
	return true
}

func (s *State) MarkGuess(playerName, guessPlayerName string, correct bool) bool {
	if s.Phase != phaseGuessing || playerName == guessPlayerName {
		return false
	}

	_, ok := s.Players[playerName]
	if !ok || s.DrawingPlayer != playerName {
		return false
	}

	guessPlayer, ok := s.Players[guessPlayerName]
	if !ok {
		return false
	}

	if correct {
		s.CorrectPlayer = guessPlayerName
		s.CorrectAnswer = guessPlayer.Guess
		guessPlayer.Points++
		s.Phase = phaseResult
	}

	guessPlayer.GuessIncorrect = !correct

	return true
}

func (s *State) EndRound() bool {
	if s.Phase != phaseResult {
		return false
	}

	s.Image = ""
	s.CorrectAnswer = ""
	s.nextDrawer()
	return true
}

func (s *State) nextDrawer() {
	if len(s.Players) == 0 {
		return
	}

	var (
		names = []string{}
	)

	for name, player := range s.Players {
		names = append(names, name)

		player.Guess = ""
		player.GuessIncorrect = false
	}

	sort.Strings(names)

	i := sort.SearchStrings(names, s.DrawingPlayer) + 1
	s.DrawingPlayer = names[i%len(names)]

	s.Phase = phaseDrawing
}
