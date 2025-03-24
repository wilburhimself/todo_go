package session

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	SessionCookieName = "todo_session"
	SessionDuration   = 24 * time.Hour
)

type Session struct {
	UserID    uint
	CreatedAt time.Time
}

var (
	sessions     = make(map[string]Session)
	sessionMutex = &sync.Mutex{}
)

func GenerateSessionID() string {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	sessionID := uuid.NewString()
	sessions[sessionID] = Session{
		CreatedAt: time.Now(),
	}

	return sessionID
}

func CreateSession(w http.ResponseWriter, userID uint) string {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	sessionID := uuid.NewString()
	sessions[sessionID] = Session{
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	cookie := http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(SessionDuration.Seconds()),
	}
	http.SetCookie(w, &cookie)

	return sessionID
}

func GetSession(r *http.Request) (Session, bool) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return Session{}, false
	}

	session, exists := sessions[cookie.Value]
	if !exists || time.Since(session.CreatedAt) > SessionDuration {
		// Session expired or doesn't exist
		return Session{}, false
	}

	return session, true
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	// Get the session cookie
	cookie, err := r.Cookie(SessionCookieName)
	if err == nil {
		// Remove session from map if it exists
		delete(sessions, cookie.Value)
	}

	// Clear the cookie
	expiredCookie := http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, &expiredCookie)
}
