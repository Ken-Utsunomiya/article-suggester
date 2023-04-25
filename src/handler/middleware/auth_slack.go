package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func AuthSlack(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
		verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(io.TeeReader(r.Body, &verifier))

		if err = verifier.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
