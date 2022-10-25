package middlewares

import (
	"context"
	"go-server/cmd/api/session"

	"net/http"
)

var UserSessions *session.Manager

// initialize in init() function
func init() {
	UserSessions, _ = session.NewManager("memory", "sid", 3600)
	go UserSessions.GC()
}

type Middleware func(http.Handler) http.Handler

func Chain(n http.Handler, m []Middleware) http.Handler {
	if len(m) < 1 {
		return n
	}
	wrapped := n

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

var ShareMdware = []Middleware{
	SessionMiddleware,
	EnableCORS,
}

func EnableCORS(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		n.ServeHTTP(w, r)
	})
}

type CtxKey string

func SessionMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Load session if exists or create new session
		session := UserSessions.SessionStart(w, r)
		sid := session.SessionID()
		// set session id in context
		r = r.Clone(context.WithValue(r.Context(), CtxKey("sid"), sid))
		n.ServeHTTP(w, r)
	})
}
