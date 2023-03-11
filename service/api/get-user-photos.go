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
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'id dell'utente di cui voglio ottenere le foto
	reqUser, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	userId, _ := strconv.Atoi(r.Header.Get("Authorization"))
	ris, phList := rt.db.GetUserPhotos(userId, reqUser)

	switch ris {

	case 0:

		var photoList []CompletePost
		var post CompletePost
		for x := 0; x < len(phList); x++ {
			post.ID = phList[x].ID
			post.USERNAME = phList[x].USERNAME
			post.LIKES = phList[x].LIKES
			post.COMMENTS = phList[x].COMMENTS
			post.DATE = phList[x].DATE
			photoList = append(photoList, post)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(photoList)

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
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
	}

}
