package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// rt.router.GET("/users/:userId/profile", rt.wrap(rt.getUserProfile))

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")
	authI, _ := strconv.Atoi(auth)

	// Prendo il cod utente indicato nel path
	reqUser := strings.Split(r.RequestURI, "/")[2]

	reqUserI, _ := strconv.Atoi(reqUser)

	ris, userProfile := rt.db.GetUserProfile(authI, reqUserI)

	switch ris {
	case 0:
		// user, _ := json.Marshal(userProfile)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(userProfile)

	case -1:
		ctx.Logger.Error("User not exist")
		http.Error(w, "User not exist", http.StatusBadRequest)

	case -2:
		ctx.Logger.Error("User you want to know the information  does not exist")
		http.Error(w, "User you want to know the information  does not exist", http.StatusBadRequest)

	case -3:
		ctx.Logger.Error("User you want to know the infromation has banned you")
		http.Error(w, "User you want to know the infromation has banned you", http.StatusForbidden)

	case -4:
		ctx.Logger.Error("You ban the user you want to know the information")
		http.Error(w, "You ban the user you want to know the information", http.StatusMethodNotAllowed)

	case -5:
		ctx.Logger.Error("Error during execution")
		http.Error(w, "Error during execution", http.StatusInternalServerError)

	}
}
