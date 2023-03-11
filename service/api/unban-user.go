package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// VA BENE
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]
	banId := strings.Split(r.RequestURI, "/")[4]

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == reqUser && auth != banId {

		reqUser, _ := strconv.Atoi(reqUser)
		banId, _ := strconv.Atoi(banId)

		ris, username := rt.db.UnBanUser(reqUser, banId)

		switch ris {

		case 0:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(username)

		case -1:
			ctx.Logger.Error(UserIdNotFound)
			http.Error(w, UserIdNotFound, http.StatusBadRequest)

		case -2:
			ctx.Logger.Error(UserId2NotFound)
			http.Error(w, UserId2NotFound, http.StatusNotFound)

		case -3:
			ctx.Logger.Error(UserIdBanned)
			http.Error(w, UserIdBanned, http.StatusForbidden)

		case -4:
			ctx.Logger.Error(userNotBanned)
			http.Error(w, userNotBanned, http.StatusNotAcceptable)

		case -5:
			ctx.Logger.Error(ErrorServerExecution)
			http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		}

	} else {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
	}
}
