package session

import (
	"github.com/gin-contrib/sessions/cookie"
)

var SessionStore = cookie.NewStore([]byte("ginframe"))
