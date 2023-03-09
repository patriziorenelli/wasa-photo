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
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]
	// Prendo il cod dell'utente da seguire
	followId := strings.Split(r.RequestURI, "/")[4]

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == reqUser && auth != followId {

		// Se si sta cercando di seguire se stessi si ritorna errore
		if auth == followId {
			ctx.Logger.Error(FollowHimself)
			w.WriteHeader(http.StatusUnauthorized)
		}

		reqUser, _ := strconv.Atoi(reqUser)
		followId, _ := strconv.Atoi(followId)

		ris, username := rt.db.FollowUser(reqUser, followId)

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
			ctx.Logger.Error("You already follow the user")
			w.WriteHeader(http.StatusUnauthorized)

		case -6:
			ctx.Logger.Error(ErrorServerExecution)
			http.Error(w, ErrorServerExecution, http.StatusInternalServerError)

		}

	} else {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
	}

}
