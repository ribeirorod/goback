package session

// 1. Create a Unique Session ID
// 2. Open Data Storage Space for Session : Memory or DB
// 3. Send a Unique session ID to the client
// Either use the response line, header or body to accomplish this

// Deal with expired sessions

// Cookies: the server can easily use Set-cookie inside of a response header to send a session id to a client,
// and a client can then use this cookie for future requests;
// we often set the expiry time for cookies containing session information to 0,
// which means the cookie will be saved in memory and only deleted after users have close their browsers.

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Session Management

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxlifetime int64)
}

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}

var provides = make(map[string]Provider)

// NewManager init session manager
func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil, it panics
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

// Session ID
func (manager *Manager) SessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionInit init session
func (manager *Manager) SessionInit(sid string) (Session, error) {
	session, err := manager.provider.SessionInit(sid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// SessionRead read session by sid
func (manager *Manager) SessionRead(sid string) (Session, error) {
	session, err := manager.provider.SessionRead(sid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// SessionDestroy delete session by id
func (manager *Manager) SessionDestroy(sid string) error {
	err := manager.provider.SessionDestroy(sid)
	return err
}

// GC start gc routine
func (manager *Manager) GC() {
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime)*time.Second, func() { manager.GC() })
}

// SessionStart start session by sid
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.SessionID()
		session, _ = manager.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)

	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.SessionRead(sid)
	}
	return
}

//
