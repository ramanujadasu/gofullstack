package handlers

import (
	"net/http"

	"github.com/ramanujadasu/gofullstack/volume3/section5/gopherface/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
