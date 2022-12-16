package api

import (
	"encoding/json"
	"net/http"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// 
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

}
