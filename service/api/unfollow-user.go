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
			_ = json.NewEncoder(w).Encode(username)

		case -1:
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -2:
			ctx.Logger.Error("User you want to unfollow does not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -3:
			ctx.Logger.Error("User you want to unfollow has banned you")
			w.WriteHeader(http.StatusUnauthorized)

		case -4:
			ctx.Logger.Error("You ban the user you want to unfollow")
			w.WriteHeader(http.StatusUnauthorized)

		case -5:
			ctx.Logger.Error("You don't already follow the user")
			w.WriteHeader(http.StatusUnauthorized)

		case -6:
			ctx.Logger.Error("Error during execution")
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		ctx.Logger.Error(Fail_Auth)
		w.WriteHeader(http.StatusUnauthorized)
	}
}
