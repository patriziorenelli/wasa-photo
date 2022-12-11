package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// PRENDE CORRETTAMENTE L'USERNAME PASSATO MA DA ERRORE SUL CONTROLLO VALIDITA' -> E' SBAGLIATA LA FUNZIONE DI VALIDAZIONE

func (rt *_router) DoLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// deve ritornare un json contentente id dell'utente
	var user Username
	var us User
	err := json.NewDecoder(r.Body).Decode(&user)

	// controllo che l'username passato sia nel formato corretto
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.UsernameIsValid() {
		// entra qui anche se dovrebbe essere valido
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbUser, err := rt.db.DoLogin(user.UsernameToDatabase())

	// se dbUser sarebbe -1 non esiste l'utente
	if err != nil {
		// Utente non esiste
		// bisogna chiamare la funzione che si occupa di creare il nuovo utente

		us.FromDatabase(dbUser)
		us.ID = -1
		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(us)

		return
	}

	// qui forse andrebbe fatto qualcosa per la sicurezza
	var id UserId
	id.ID = dbUser.ID
	id.FromUserDatabase(dbUser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(us)

}
