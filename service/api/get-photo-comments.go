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
func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")

	// Prendo l'id del post di cui vogliamo sapere i commenti
	postId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	userId, _ := strconv.Atoi(auth)
	ris, commentList := rt.db.GetPhotoComment(userId, postId)

	switch ris {

	case 0:
		var arC []string
		for x := 0; x < len(commentList); x++ {
			//user := UserId{  followers[x] }
			//v, _ := json.Marshal(user)
			//fmt.Println(string(v))
			v, _ := json.Marshal(commentList[x])
			arC = append(arC, string(v))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arC)

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -2:
		ctx.Logger.Error("Photo  does not exist")
		w.WriteHeader(http.StatusUnauthorized)

	case -3:
		ctx.Logger.Error("User who posted the photo whose comment you want to know has blocked you")
		w.WriteHeader(http.StatusUnauthorized)

	case -4:
		ctx.Logger.Error("You ban the user that posted the photo")
		w.WriteHeader(http.StatusUnauthorized)

	case -5:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)
	}

}
