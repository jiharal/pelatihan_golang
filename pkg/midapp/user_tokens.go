package midapp

import (
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("sessions_app"))

func GetSession() *sessions.CookieStore {
	return store
}
