package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// TESTARE LA FUNZIONALITA'

// curl -X PUT  http://localhost:3000/users/1/username -H "Authorization: 1" -H "Content-Type: application/json" -d '{"username": "marione_12"}'
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

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Converto l'id utente in un int
		userId, _ := strconv.Atoi(reqUser)

		dbUser, err := rt.db.SetMyUserName(userId, user.USERNAME)

		if err != nil {
			ctx.Logger.WithError(err).Error("Error during change username")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(dbUser)

		}

	} else {
		ctx.Logger.Error("User aren't authorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
