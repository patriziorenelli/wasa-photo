package api

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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

	if rt.db.CheckBan(authId, userId) == 0 {
		ctx.Logger.Error("You ban the photo owner")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if rt.db.CheckBan(userId, authId) == 0 {
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

	mydir, erro := os.Getwd()
	if erro != nil {
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	path := mydir + "/photos/" + strconv.Itoa(userId) + "/" + strconv.Itoa(postId) + ".jpg"

	file, erOpen := os.Open(path)

	if erOpen != nil {

		ctx.Logger.Error("Photo file not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//  CONVERSIONE E SCRITTURA RITORNO

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	var photoFile PhotoFile

	photoFile.PHOTOFILE = encoded

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoFile)

}
