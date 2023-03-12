package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// Va bene
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]
	unfollowId := strings.Split(r.RequestURI, "/")[4]

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == reqUser && auth != unfollowId {

		reqUser, _ := strconv.Atoi(reqUser)
		unfollowId, _ := strconv.Atoi(unfollowId)

		ris, username := rt.db.UnfollowUser(reqUser, unfollowId)

		switch ris {

		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
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
			ctx.Logger.Error(userId2Banned)
			http.Error(w, userId2Banned, http.StatusMethodNotAllowed)

		case -5:
			ctx.Logger.Error(UnfollowedUser)
			http.Error(w, UnfollowedUser, http.StatusPaymentRequired)

		case -6:
			ctx.Logger.Error(ErrorServerExecution)
			http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		}
	} else {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
	}
}
