package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// DA FARE
// /users/:userId/stream"
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth, _ := strconv.Atoi(r.Header.Get("Authorization"))

	reqUser, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])

	// Se l'utente è autorizzato a visualizzare lo stream dell'utente richiesto
	if auth == reqUser {

		// Estraiamo dall'url i query parameters limit e startIndex
		var err error
		err = nil
		queryPar := r.URL.Query()
		var limit int
		var startIndex int

		if (queryPar).Get("limit") != "" {
			limit, err = strconv.Atoi(queryPar.Get("limit"))
		} else {
			limit = 10
		}

		if (queryPar).Get("startIndex") != "" {
			startIndex, err = strconv.Atoi(queryPar.Get("startIndex"))
		} else {
			startIndex = 0
		}

		if err != nil {
			startIndex = 0
			limit = 10
		}

		ris, phList := rt.db.GetMyStream(auth, limit, startIndex)

		// []CompletePost
		switch ris {

		case 0:
			var photoList []CompletePost
			var post CompletePost
			for x := 0; x < len(phList); x++ {
				post.ID = phList[x].ID
				post.USERID = phList[x].USERID
				post.USERNAME = phList[x].USERNAME
				post.LIKES = phList[x].LIKES
				post.COMMENTS = phList[x].COMMENTS
				post.DATE = phList[x].DATE
				photoList = append(photoList, post)
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(photoList)

		case -1:
			ctx.Logger.Error("User not exist")
			w.WriteHeader(http.StatusUnauthorized)

		case -2:
			ctx.Logger.Error("Error during execution")
			w.WriteHeader(http.StatusUnauthorized)
		}

	} else {
		ctx.Logger.Error(Fail_Auth)
		w.WriteHeader(http.StatusUnauthorized)
	}

}
