package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace
	userId := strings.Split(r.RequestURI, "/")[2]

	// Controllo che l'autenticazione vada a buon fine
	if auth != userId {
		ctx.Logger.Error("Failed authentication")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		ctx.Logger.Error("Error during image check")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	mimeType := http.DetectContentType(bytes)

	if !strings.HasPrefix(mimeType, "image/") {
		ctx.Logger.Error("File is not a valid image")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usId, _ := strconv.Atoi(userId)

	ris, photoId := rt.db.UploadPhoto(usId)

	switch ris {

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)
		return
	case -2:
		ctx.Logger.Error("Error during saving into database")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	path := "/photos/" + userId + "/" + strconv.Itoa(photoId) + ".jpg"

	// Creo la directory, in caso di errore elimino anche il record relativo alla nuova foto
	if os.MkdirAll(filepath.Dir(path), os.ModePerm) != nil {
		ctx.Logger.Error("Error during directory creation")
		w.WriteHeader(http.StatusUnauthorized)
		_ = rt.db.DeletePhotoRecord(photoId)
		return
	}

	// Salvo l'immagine, in caso di errore elimino il record relativo alla nuova foto
	if err = os.WriteFile(path, bytes, 0644); err != nil {
		ctx.Logger.Error("Error save photo file")
		w.WriteHeader(http.StatusUnauthorized)
		_ = rt.db.DeletePhotoRecord(photoId)
		return
	}

	ctx.Logger.Error("Photo uploaded")
	w.WriteHeader(http.StatusUnauthorized)
	return
}
