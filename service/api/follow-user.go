package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"encoding/json"
    "strconv"
	
)

// 	rt.router.PUT("/users/:userId/followUser/:userId2", rt.wrap(rt.followUser))
// VA BENE 
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]
	followId := strings.Split(r.RequestURI, "/")[4]

	

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == reqUser && auth != followId {
		
		reqUser, _ := strconv.Atoi(reqUser)
		followId, _ := strconv.Atoi(followId)


		ris, username := rt.db.FollowUser(reqUser, followId)
		// FARE TEST SUL CORRETTO FUNZIONAMENTO
		// Qui bisogna controllare il res e mandare il json adeguato se serve 
		switch ris{

		case 0:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(username)
			return 

		case -1:
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -2:
			ctx.Logger.Error("User you want to follow does not exist")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -3:
			ctx.Logger.Error("User you want to follow has banned you")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -4:
			ctx.Logger.Error("You ban the user you want to follow")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -5:
			ctx.Logger.Error("Already followed the user")
			w.WriteHeader(http.StatusUnauthorized)
			return

		case -6:
			ctx.Logger.Error("Error during execution")
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

	} else {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
