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
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'id dell'utente di cui voglio ottenere le foto
	reqUser, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	userId, _ := strconv.Atoi(r.Header.Get("Authorization"))
	ris, phList := rt.db.GetUserPhotos(userId, reqUser)

	switch ris {

	case 0:
		var photoList []string
		for x := 0; x < len(phList); x++ {
			//user := UserId{  followers[x] }
			//v, _ := json.Marshal(user)
			//fmt.Println(string(v))
			v, _ := json.Marshal(phList[x])
			photoList = append(photoList, string(v))
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photoList)

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -2:
		ctx.Logger.Error("User whose photos you want does not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -3:
		ctx.Logger.Error("User whose photos you want has blocked you")
		w.WriteHeader(http.StatusUnauthorized)

	case -4:
		ctx.Logger.Error("You ban the other user")
		w.WriteHeader(http.StatusUnauthorized)

	case -5:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)
	}

}
