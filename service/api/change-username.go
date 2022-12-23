package api

import (
	// "encoding/json"
	"fmt"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	// "reflect"
	"strings"
	"encoding/json"
	"strconv"
)

//  curl -X PUT  http://localhost:3000/users/1/username -H "Authorization: 1" -H "Content-Type: application/json" -d '{"username": "marione_12"}'

// CAPIRE COME AGGIUNGERE L'AUTH  -> va bene usare quest curl:
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// r.Header.Get("Authorization")  ->  Prendo l'autenticazione
	x := r.Header.Get("Authorization")
	// fmt.Print(x)



	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]

	fmt.Print(x)

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

		userId, _ := strconv.Atoi(reqUser)

		dbUser, err := rt.db.SetMyUserName(userId, user.USERNAME)
		// qui bisogna fare il json da ritornare 
		fmt.Print(dbUser)

		/* Qui bisogna chiamare la funzione per fare il change username e poi ritornare
		{
			"newUserName": "Marco12"
		}
		*/

	} else {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	/*
	// deve ritornare un json contentente id dell'utente
		var user Username
		var us User
		err := json.NewDecoder(r.Body).Decode(&user)

		// controllo che l'username passato sia nel formato corretto
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if !user.UsernameIsValid() {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dbUser, err := rt.db.DoLogin(user.UsernameToDatabase())

		// se dbUser sarebbe -1 non esiste l'utente
		if err == nil && dbUser.ID == -2 {
			// Utente non esiste
			// bisogna chiamare la funzione che si occupa di creare il nuovo utente

			newUser, err := rt.db.CreateUser(dbUser.USERNAME)


			if err != nil && dbUser.ID == -1{
				ctx.Logger.WithError(err).Error("Error during creation user")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}else if err != nil && dbUser.ID == -2{
				ctx.Logger.WithError(err).Error("Error during extract new userId")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			us.ID = newUser.ID
			us.USERNAME = newUser.USERNAME

		}else if err != nil && dbUser.ID == -1{
			ctx.Logger.WithError(err).Error("Error during find userId")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}else{
			us.ID = dbUser.ID
			us.USERNAME = user.USERNAME
		}

		// qui forse andrebbe fatto qualcosa per la sicurezza



		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(us)

	*/

}
