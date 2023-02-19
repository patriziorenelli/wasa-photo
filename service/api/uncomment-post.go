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
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")
	// Prendo l'id del post a cui togliere il commento
	phId := strings.Split(r.RequestURI, "/")[2]

	// Prendo l'id del commento da eliminare
	cmId := strings.Split(r.RequestURI, "/")[4]

	userId, _ := strconv.Atoi(auth)
	postId, _ := strconv.Atoi(phId)
	commentId, _ := strconv.Atoi(cmId)

	ris := rt.db.UncommentPhoto(userId, postId, commentId)

	switch ris {

	case 0:
		var risultato Result
		risultato.TEXT = Done
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(risultato)

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -2:
		ctx.Logger.Error("The post doesn't exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -3:
		ctx.Logger.Error("You banned other user")
		w.WriteHeader(http.StatusUnauthorized)

	case -4:
		ctx.Logger.Error("The other user blocked you")
		w.WriteHeader(http.StatusUnauthorized)

	case -5:
		ctx.Logger.Error("Comment doesn't exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -6:
		ctx.Logger.Error("Failed authentication")
		w.WriteHeader(http.StatusUnauthorized)

	case -7:
		ctx.Logger.Error("Comment not associated with the post")
		w.WriteHeader(http.StatusUnauthorized)

	case -8:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)

	}
}
