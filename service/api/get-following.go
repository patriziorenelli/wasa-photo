package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// VA BENE
func (rt *_router) getUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")
	authI, _ := strconv.Atoi(auth)

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]

	reqUserI, _ := strconv.Atoi(reqUser)

	ris, following := rt.db.GetUserFollowing(authI, reqUserI)

	switch ris {

	case 0:

		// Ritorno la lista degli utenti che  l'utente segue
		var arF []string
		for x := 0; x < len(following); x++ {
			//user := UserId{  followers[x] }
			//v, _ := json.Marshal(user)
			//fmt.Println(string(v))
			v, _ := json.Marshal(following[x])
			arF = append(arF, string(v))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arF)

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -2:
		ctx.Logger.Error("User you want to know the following  does not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -3:
		ctx.Logger.Error("User you want to know the following has banned you")
		w.WriteHeader(http.StatusUnauthorized)

	case -4:
		ctx.Logger.Error("You ban the user you want to know the following")
		w.WriteHeader(http.StatusUnauthorized)

	case -5:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)

	}

}
