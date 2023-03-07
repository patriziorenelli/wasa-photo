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
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace
	postId := strings.Split(r.RequestURI, "/")[2]
	// Prendo il cod utente indicato nel path
	userId := strings.Split(r.RequestURI, "/")[4]

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == userId {

		userId, _ := strconv.Atoi(userId)
		postId, _ := strconv.Atoi(postId)

		ris := rt.db.LikePhoto(userId, postId)

		switch ris {

		case 0:
			var risultato Result
			risultato.TEXT = Done
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(risultato)

		case -1:
			ctx.Logger.Error(UserIdNotFound)
			http.Error(w, UserIdNotFound, http.StatusBadRequest)

		case -2:
			ctx.Logger.Error(photoNotFound)
			http.Error(w, photoNotFound, http.StatusProxyAuthRequired)

		case -3:
			ctx.Logger.Error(userId2Banned)
			http.Error(w, userId2Banned, http.StatusMethodNotAllowed)

		case -4:
			ctx.Logger.Error(UserIdBanned)
			http.Error(w, UserIdBanned, http.StatusForbidden)

		case -5:
			ctx.Logger.Error("You already liked the post")
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
