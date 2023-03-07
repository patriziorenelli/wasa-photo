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
		http.Error(w, "Username not valid", http.StatusLengthRequired)
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
		ctx.Logger.Error("Username doesn't exist")
		http.Error(w, "Username doesn't exist", http.StatusNotFound)
	case -2:
		ctx.Logger.Error("UserId doesn't exist")
		http.Error(w, "UserId doesn't exist", http.StatusBadRequest)
	case -3:
		ctx.Logger.Error("Error during execution")
		http.Error(w, "Error during execution", http.StatusInternalServerError)
		
	}

}
