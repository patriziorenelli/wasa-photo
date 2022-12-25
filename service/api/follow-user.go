package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	// 	"encoding/json"
    "strconv"
	"fmt"
)

// 	rt.router.PUT("/users/:userId/followUser/:userId2", rt.wrap(rt.followUser))

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]
	followId := strings.Split(r.RequestURI, "/")[4]

	

	// Se l'autenticazione va a buon fine e si sta cercando di seguire un altro user, si invia la richiesta di follow
	if auth == reqUser && auth != followId {
		
		reqUser, _ := strconv.Atoi(reqUser)
		followId, _ := strconv.Atoi(followId)


		res, err := rt.db.FollowUser(reqUser, followId)
		// FARE TEST SUL CORRETTO FUNZIONAMENTO
		// Qui bisogna controllare il res e mandare il json adeguato se serve 
		fmt.Print("\n FINITO")
		fmt.Print(res, err)

	} else {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
