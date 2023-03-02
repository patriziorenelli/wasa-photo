package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// DA FARE
func (rt *_router) getUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// deve ritornare un json contentente id dell'utente
	var userId UserId

	auth, _ := strconv.Atoi(r.Header.Get("Authorization"))

	var user Username
	err := json.NewDecoder(r.Body).Decode(&user)

	// controllo che l'username passato sia nel formato corretto
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.UsernameIsValid() {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbUser, errId := rt.db.GetUsreId(auth, user.UsernameToDatabase())

	switch errId {

	case 0:
		userId.USERID = dbUser.USERID
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(userId)
	case -1:
		ctx.Logger.WithError(err).Error("Username doesn't exist")
		w.WriteHeader(http.StatusInternalServerError)
	case -2:
		ctx.Logger.WithError(err).Error("UserId doesn't exist")
		w.WriteHeader(http.StatusInternalServerError)
	case -3:
		ctx.Logger.WithError(err).Error("Error during creation user")
		w.WriteHeader(http.StatusInternalServerError)

	}

}
