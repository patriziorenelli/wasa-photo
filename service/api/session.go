package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Read username from the request body.
	var user username

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.usernameIsValid() {
		// Here we validated the username structure content, but it isn't valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the user in the database. Note that this function will return a new instance of the
	dbfountain, err := rt.db.createUser(User.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the fountain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	user.FromDatabase(dbfountain)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(fountain)
}
