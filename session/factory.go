package session

import "sync"

var sessions = map[string]*Session{}
var m sync.Mutex

func GetSession(sId string) (s *Session, ok bool) {
	m.Lock()
	defer m.Unlock()

	s, ok = sessions[sId]

	return
}

func GetSessions() map[string]*Session {
	return sessions
}

func AddSession(sId string, s *Session) {
	m.Lock()
	defer m.Unlock()

	sessions[sId] = s
}

func RemoveSession(sId string) {
	m.Lock()
	defer m.Unlock()

	delete(sessions, sId)
}
