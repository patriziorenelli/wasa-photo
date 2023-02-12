package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

func (rt *_router) getPhotoLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace
	postId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	userId, _ := strconv.Atoi(auth)
	ris, userLike := rt.db.GetPhotoLike(userId, postId)

	switch ris {

	case 0:
		var arL []string
		for x := 0; x < len(userLike); x++ {
			//user := UserId{  followers[x] }
			//v, _ := json.Marshal(user)
			//fmt.Println(string(v))
			v, _ := json.Marshal(userLike[x])
			arL = append(arL, string(v))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arL)

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -2:
		ctx.Logger.Error("Photo  does not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -3:
		ctx.Logger.Error("User who posted the photo whose number of likes you want to know has blocked you")
		w.WriteHeader(http.StatusUnauthorized)

	case -4:
		ctx.Logger.Error("You ban the user that posted the photo")
		w.WriteHeader(http.StatusUnauthorized)

	case -5:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)
	}

}
