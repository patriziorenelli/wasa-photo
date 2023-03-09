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
func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")
	authI, _ := strconv.Atoi(auth)

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]

	reqUserI, _ := strconv.Atoi(reqUser)

	ris, followers := rt.db.GetUserFollowers(authI, reqUserI)

	switch ris {

	case 0:
		var user UserId
		var arF []UserId
		for x := 0; x < len(followers); x++ {
			user.USERID = followers[x].USERID
			arF = append(arF, user)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arF)

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
