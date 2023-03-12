package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)
// Va bene 
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo l'id dell'utente che vuole pubblicare la foto
	userId := strings.Split(r.RequestURI, "/")[2]

	// Controllo che l'autenticazione vada a buon fine
	if auth != userId {
		ctx.Logger.Error(Fail_Auth)
		http.Error(w, Fail_Auth, http.StatusBadGateway)
		return
	}

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		ctx.Logger.Error("Error during image check")
		http.Error(w, InvalidFormat, http.StatusGone)
		return
	}

	mimeType := http.DetectContentType(bytes)
	// Controllo il tipo del file caricato
	if !strings.HasPrefix(mimeType, "image/") {
		ctx.Logger.Error(InvalidFormat)
		http.Error(w, InvalidFormat, http.StatusGone)
		return
	}

	usId, _ := strconv.Atoi(userId)

	ris, photoId := rt.db.UploadPhoto(usId)

	if ris == -1 {
		ctx.Logger.Error(UserIdNotFound)
		http.Error(w, UserIdNotFound, http.StatusBadRequest)
		return
	} else if ris == -2 {
		ctx.Logger.Error("Error during saving into database")
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		return
	}

	mydir, err := os.Getwd()
	if err != nil {
		ctx.Logger.Error("Error during directory creation")
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		_ = rt.db.DeletePhotoRecord(photoId)
		return
	}

	path := mydir + "/photos/" + userId + "/" + strconv.Itoa(photoId) + ".jpg"
	// Creo la directory, in caso di errore elimino anche il record relativo alla nuova foto
	err = os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		_ = rt.db.DeletePhotoRecord(photoId)
		return
	}

	// Salvo l'immagine, in caso di errore elimino il record relativo alla nuova foto
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
		_ = rt.db.DeletePhotoRecord(photoId)
		return
	}

	var risultato Result
	risultato.TEXT = Done
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(risultato)

}
