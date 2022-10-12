package middlewares

import (
	"go-server/cmd/api/session"
	_ "go-server/cmd/api/sessionMemory" // import memory session executes init() function
	"log"
	"net/http"
)

var userSessions *session.Manager

// initialize in init() function
func init() {
	userSessions, _ = session.NewManager("memory", "user_session", 3600)
	go userSessions.GC()
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

func SessionMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Load session if exists or create new session
		session := userSessions.SessionStart(w, r)
		sid := session.SessionID()
		log.Println("Starting session", session, sid)
		// End User session
		defer userSessions.SessionDestroy(sid)
		n.ServeHTTP(w, r)

	})
}
