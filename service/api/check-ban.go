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
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -2:
			ctx.Logger.Error("User you want to ban does not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -3:
			ctx.Logger.Error("User you want to ban has banned you")
			w.WriteHeader(http.StatusUnauthorized)

		case -4:
			ctx.Logger.Error("You haven't ban the other user")
			w.WriteHeader(http.StatusUnauthorized)

		}
	} else {
		ctx.Logger.Error(Fail_Auth)
		w.WriteHeader(http.StatusUnauthorized)
	}

}
