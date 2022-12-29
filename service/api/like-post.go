package api

import (
	// "encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// DA FARE 
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace 
	postId := strings.Split(r.RequestURI, "/")[2]
	// Prendo il cod utente indicato nel path
	userId := strings.Split(r.RequestURI, "/")[4]

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == userId  {

		userId, _ := strconv.Atoi(userId)
		postId, _ := strconv.Atoi(postId)

		ris := rt.db.LikePost(userId, postId)

		switch ris {

		case 0:
			w.Header().Set("Content-Type", "application/json")
			// _ = json.NewEncoder(w).Encode(username)
			return

		case -1:
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -2:
			ctx.Logger.Error("The post doesn't exist")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -3:
			ctx.Logger.Error("You banned other user")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -4:
			ctx.Logger.Error("The other user blocked you")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -5:
			ctx.Logger.Error("You already liked the post")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -6:
			ctx.Logger.Error("Error during execution")
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

	} else {
		ctx.Logger.Error("Failed authentication")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
