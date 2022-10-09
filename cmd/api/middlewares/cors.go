package middlewares

import (
	_ "go-server/cmd/api/memory"
	"go-server/cmd/api/session"
	"net/http"
	"regexp"
)

var globalSessions *session.Manager

// initialize in init() function
func init() {
	globalSessions, _ = session.NewManager("memory", "sid", 3600)
	go globalSessions.GC()
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		n.ServeHTTP(w, r)
	})
}

func SessionMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get Cookies from request
		//cookie := r.Header.Get("Cookie")
		if cookie, nok := r.Cookie("sid"); nok == nil {

			re, _ := regexp.Compile(`sid=([\d\w]+)[^ec;]`)
			sid := re.FindStringSubmatch(cookie.Value)[1]
			_, err := globalSessions.SessionInit(sid)

			if err != nil {
				panic(err)
			}
			defer globalSessions.SessionDestroy(sid)

		} else {

			// create session
			session := globalSessions.SessionStart(w, r)
			sid := session.SessionID()
			// set cookie
			cookie := &http.Cookie{
				Name:     "sid",
				Value:    sid,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Path:     "/",
				MaxAge:   3600}

			http.SetCookie(w, cookie)
		}

		n.ServeHTTP(w, r)

	})
}
