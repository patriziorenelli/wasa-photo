package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// Va bene 
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]

	// Controllo che l'utente sia correttamente loggato
	if auth == reqUser {
		// Ottengo il nuovo username che l'utente vuole impostare
		var user Username
		err := json.NewDecoder(r.Body).Decode(&user)
		// Controllo che l'username sia valido
		if err != nil {
			ctx.Logger.Error(ErrorServerExecution)
			http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
			return
		} else if !user.UsernameIsValid() {
			ctx.Logger.Error(UsernameNotValid)
			http.Error(w, UsernameNotValid, http.StatusLengthRequired)
			return
		}
		// Converto l'id utente in un int
		userId, _ := strconv.Atoi(reqUser)

		ris := rt.db.SetMyUserName(userId, user.USERNAME)

		switch ris {

		case 0:
			var username Username
			username.USERNAME = user.USERNAME
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(username)

		case -1:
			ctx.Logger.Error(UserIdNotFound)
			http.Error(w, UserIdNotFound, http.StatusBadRequest)

		case -2:
			ctx.Logger.Error("Username already used")
			http.Error(w, UsernameNotValid, http.StatusLengthRequired)
		case -3:
			ctx.Logger.Error(ErrorServerExecution)
			http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		}

	} else {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
	}
}
