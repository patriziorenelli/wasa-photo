package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Va bene
func (rt *_router) getUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// deve ritornare un json contentente id dell'utente
	var userId UserId

	auth, _ := strconv.Atoi(r.Header.Get("Authorization"))

	var user Username

	user.USERNAME = (r.URL.Query()).Get("username")

	if !user.UsernameIsValid() {
		ctx.Logger.Error(UsernameNotValid)
		http.Error(w, UsernameNotValid, http.StatusLengthRequired)
		return
	}

	dbUser, errId := rt.db.GetUserId(auth, user.UsernameToDatabase())

	switch errId {

	case 0:
		userId.USERID = dbUser.USERID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(userId)
	case -1:
		ctx.Logger.Error(UsernameNotFound)
		http.Error(w, UsernameNotFound, http.StatusNotFound)
	case -2:
		ctx.Logger.Error(UserIdNotFound)
		http.Error(w, UserIdNotFound, http.StatusBadRequest)
	case -3:
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)

	}

}
