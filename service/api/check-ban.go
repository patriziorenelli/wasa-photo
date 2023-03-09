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
func (rt *_router) checkUserBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	reqUser := strings.Split(r.RequestURI, "/")[2]
	banId := strings.Split(r.RequestURI, "/")[4]

	if auth == reqUser && auth != banId {

		reqUser, _ := strconv.Atoi(reqUser)
		banId, _ := strconv.Atoi(banId)

		ris, userId := rt.db.CheckUserBan(reqUser, banId)

		switch ris {

		case 0:
			// Qui devo ritornare l'id utente verificato
			var usId UserId
			usId.USERID = userId.USERID
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(usId)

		case -1:
			ctx.Logger.Error(UserIdNotFound)
			http.Error(w, UserIdNotFound, http.StatusBadRequest)

		case -2:
			ctx.Logger.Error(UserId2NotFound)
			http.Error(w, UserId2NotFound, http.StatusNotFound)

		case -3:
			ctx.Logger.Error(UserIdBanned)
			http.Error(w, UserIdBanned, http.StatusForbidden)

		case -4:
			ctx.Logger.Error("You didn't ban the user")
			http.Error(w, "You didn't ban the user", http.StatusInternalServerError)
		}
	} else {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
	}

}
