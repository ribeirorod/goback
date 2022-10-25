package session

import (
	"container/list"

	"sync"
	"time"
)

// ProverMemory memory satifies session.Provider interface.
type ProviderMemory struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // save in memory
	list     *list.List               // gc
}

func (pvdr *ProviderMemory) SessionInit(sid string) (Session, error) {
	pvdr.lock.Lock()
	defer pvdr.lock.Unlock()

	if pvdr.sessions == nil {
		pvdr.sessions = make(map[string]*list.Element, 0)
	}
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pvdr.list.PushBack(newsess)
	pvdr.sessions[sid] = element
	return newsess, nil
}

func (pder *ProviderMemory) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *ProviderMemory) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

// Session Garbage Collection - Expired Session
func (pder *ProviderMemory) SessionGC(maxlifetime int64) {
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

// Extends session life time.
func (pder *ProviderMemory) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

// SessionStore satifies session.Session interface.
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
var prvdr = &ProviderMemory{list: list.New()}

func init() {
	RegisterNewProvider("memory", prvdr)
}
