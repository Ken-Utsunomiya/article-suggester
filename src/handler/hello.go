package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
)

type SlackCommandHandler struct{}

func NewSlackCommandHandler() *SlackCommandHandler {
	return &SlackCommandHandler{}
}

func (h *SlackCommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// only accepts POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// parse the request of slash command
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch s.Command {
	case "/hello":
		params := &slack.Msg{Text: fmt.Sprintf("hello %s", s.UserName)}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	case "/qiita":
		// TODO
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
