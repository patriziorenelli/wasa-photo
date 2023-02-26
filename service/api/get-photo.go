package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])

	// Prendo l'id del post di cui vogliamo sapere i like
	postId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[4])
	authId, _ := strconv.Atoi(r.Header.Get("Authorization"))

	if rt.db.UserExist(authId) == -1 {
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if rt.db.UserExist(userId) == -1 {
		ctx.Logger.Error("The photo owner does not exist")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if rt.db.CheckBan(authId, userId) == -1 {
		ctx.Logger.Error("You ban the photo owner")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if rt.db.CheckBan(userId, authId) == -1 {
		ctx.Logger.Error("The photo owner ban you")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err := rt.db.GetPhoto(postId)
	if err == -1 {
		ctx.Logger.Error("The photo does not exist")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	path := "/photos/" + strconv.Itoa(userId) + "/" + strconv.Itoa(postId) + ".jpg"

	file, erOpen := os.Open(path)

	if erOpen != nil {
		ctx.Logger.Error("Photo file not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	defer file.Close()

	_, erOpen = io.Copy(w, file)

	if erOpen != nil {
		ctx.Logger.Error("Error copying file")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
