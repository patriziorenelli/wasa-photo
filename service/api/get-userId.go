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
		ctx.Logger.Error("Username not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbUser, errId := rt.db.GetUsreId(auth, user.UsernameToDatabase())

	switch errId {

	case 0:
		userId.USERID = dbUser.USERID
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(userId)
	case -1:
		ctx.Logger.Error("Username doesn't exist")
		w.WriteHeader(http.StatusInternalServerError)
	case -2:
		ctx.Logger.Error("UserId doesn't exist")
		w.WriteHeader(http.StatusInternalServerError)
	case -3:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusInternalServerError)

	}

}
