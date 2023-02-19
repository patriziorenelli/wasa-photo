package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// DA TESTARE DI NUOVO
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
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if !user.UsernameIsValid() {
			ctx.Logger.Error("Username not valid")
			w.WriteHeader(http.StatusBadRequest)
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
			_ = json.NewEncoder(w).Encode(username)

		case -1:
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -2:
			ctx.Logger.Error("Username already used")
			w.WriteHeader(http.StatusUnauthorized)

		case -3:
			ctx.Logger.Error("Error during execution")
			w.WriteHeader(http.StatusUnauthorized)
		}

	} else {
		ctx.Logger.Error(Fail_Auth)
		w.WriteHeader(http.StatusUnauthorized)
	}
}
