package game

import "github.com/pmcatominey/sketch-game/state"

type command struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type reply struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func replyErr(err error) reply {
	return reply{
		Type: "error",
		Data: map[string]string{
			"error": err.Error(),
		},
	}
}

func replyHello(key string) reply {
	return reply{
		Type: "hello",
		Data: map[string]string{
			"key": key,
		},
	}
}

func replyState(s *state.State) reply {
	return reply{
		Type: "state",
		Data: s,
	}
}

func replyChat(player, message string) reply {
	return reply{
		Type: "chat",
		Data: map[string]string{
			"player":  player,
			"message": message,
		},
	}
}

func replyEvent(event string) reply {
	return reply{
		Type: "event",
		Data: map[string]string{
			"event": event,
		},
	}
}
