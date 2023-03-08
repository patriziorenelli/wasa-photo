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
func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo l'id del post di cui vogliamo sapere i commenti
	postId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	userId, _ := strconv.Atoi(auth)
	ris, commentList := rt.db.GetPhotoComment(userId, postId)

	switch ris {

	case 0:

		var arC []Comment
		var commento Comment
		for x := 0; x < len(commentList); x++ {
			// Creo il commento
			commento.UID = commentList[x].UID
			commento.NAME = commentList[x].NAME
			commento.TEXT = commentList[x].TEXT
			commento.CID = commentList[x].CID
			commento.DATE = commentList[x].DATE
			commento.PHID = commentList[x].PHID
			arC = append(arC, commento)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(arC)

	case -1:
		ctx.Logger.Error(UserIdNotFound)
		http.Error(w, UserIdNotFound, http.StatusBadRequest)
	case -2:
		ctx.Logger.Error(photoNotFound)
		http.Error(w, photoNotFound, http.StatusProxyAuthRequired)

	case -3:
		ctx.Logger.Error(UserIdBanned)
		http.Error(w, UserIdBanned, http.StatusForbidden)

	case -4:
		ctx.Logger.Error(userId2Banned)
		http.Error(w, userId2Banned, http.StatusMethodNotAllowed)

	case -5:
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
	}

}
