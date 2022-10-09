package memory

import (
	"container/list"
	"go-server/cmd/api/session"
	"sync"
	"time"
)

type Provider struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // save in memory
	list     *list.List               // gc
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if pder.sessions == nil {
		pder.sessions = make(map[string]*list.Element, 0)
	}
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}

		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

type SessionStore struct {
	sid          string                      // session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session store
}

// Set value in session.
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	prvdr.SessionUpdate(st.sid)
	return nil
}

// Get value from session.
func (st *SessionStore) Get(key interface{}) interface{} {
	prvdr.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

// Delete value in session.
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	prvdr.SessionUpdate(st.sid)
	return nil
}

// SessionID get session id.
func (st *SessionStore) SessionID() string {
	return st.sid
}

// SessionStore is a memory session store.
var prvdr = &Provider{list: list.New()}

func init() {
	session.Register("memory", prvdr)
}
