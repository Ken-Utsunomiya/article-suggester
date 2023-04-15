package main

import (
	"net/http"

	"github.com/projects/article-suggester/handler/router"
)

func main() {
	port := ":8080"
	mux := router.NewRouter()
	http.ListenAndServe(port, mux)
}
