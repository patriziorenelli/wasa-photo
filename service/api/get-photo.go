package api

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)
// Va bene 
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])

	// Prendo l'id del post di cui vogliamo sapere i like
	postId, _ := strconv.Atoi(strings.Split(r.RequestURI, "/")[4])
	authId, _ := strconv.Atoi(r.Header.Get("Authorization"))

	if rt.db.UserExist(authId) == -1 {
		ctx.Logger.Error(UserIdNotFound)
		http.Error(w, UserIdNotFound, http.StatusBadRequest)
		return
	}

	if rt.db.UserExist(userId) == -1 {
		ctx.Logger.Error(UserId2NotFound)
		http.Error(w, UserId2NotFound, http.StatusNotFound)
		return
	}

	if rt.db.CheckBan(authId, userId) == 0 {
		ctx.Logger.Error(userId2Banned)
		http.Error(w, userId2Banned, http.StatusMethodNotAllowed)
		return
	}

	if rt.db.CheckBan(userId, authId) == 0 {
		ctx.Logger.Error(UserIdBanned)
		http.Error(w, UserIdBanned, http.StatusForbidden)
		return
	}

	_, err := rt.db.GetPhoto(postId)
	if err == -1 {
		ctx.Logger.Error(photoNotFound)
		http.Error(w, photoNotFound, http.StatusProxyAuthRequired)
		return
	}

	mydir, erro := os.Getwd()
	if erro != nil {
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		return
	}

	path := mydir + "/photos/" + strconv.Itoa(userId) + "/" + strconv.Itoa(postId) + ".jpg"

	file, erOpen := os.Open(path)

	if erOpen != nil {
		ctx.Logger.Error("Photo file not found")
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		return
	}

	// Lettura del file
	reader := bufio.NewReader(file)
	content, _ := io.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	var photoFile PhotoFile

	photoFile.PHOTOFILE = encoded

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photoFile)

}
