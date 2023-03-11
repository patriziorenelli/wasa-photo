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
		ctx.Logger.Error(commentNotFound)
		http.Error(w, commentNotFound, http.StatusConflict)

	case -6:
		ctx.Logger.Error(UserNotAuthorized)
		http.Error(w, UserNotAuthorized, http.StatusUnauthorized)

	case -7:
		ctx.Logger.Error("Comment not associated with the photo")
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
	case -8:
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)

	}
}
