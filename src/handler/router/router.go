package router

import (
	"net/http"

	"github.com/projects/article-suggester/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/slack/command", handler.NewSlackCommandHandler().ServeHTTP)

	return mux
}
