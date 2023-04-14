package main

import (
	"net/http"
	"os"

	"github.com/projects/article-suggester/handler/router"
)

func main() {
	port := os.Getenv("PORT")
	mux := router.NewRouter()
	http.ListenAndServe(port, mux)
}
