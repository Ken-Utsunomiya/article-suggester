package router

import (
	"net/http"

	"github.com/projects/article-suggester/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handler.NewHelloHandler().ServeHTTP)

	return mux
}
