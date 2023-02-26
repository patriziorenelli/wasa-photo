package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	// "io"
	"net/http"
	"os"
	// "path/filepath"
	"strconv"
	"strings"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo l'id del proprietario del post
	userId := strings.Split(r.RequestURI, "/")[2]

	// Prendo l'id del post da eliminare
	photoId := strings.Split(r.RequestURI, "/")[4]

	if auth != userId {
		ctx.Logger.Error(Fail_Auth)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	phId, _ := strconv.Atoi(photoId)
	usId, _ := strconv.Atoi(userId)

	ris := rt.db.DeletePhoto(usId, phId)

	switch ris {
	case 0:
		// tutto ok faccio eliminazione del file
		var risultato Result
		risultato.TEXT = Done
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(risultato)
	case -1:
		// foto non esistente
		ctx.Logger.Error("The photo does not exist")
		w.WriteHeader(http.StatusUnauthorized)
	case -2:
		// utente autenticato non è il proprietario del file
		ctx.Logger.Error("The user cannot delete the photo")
		w.WriteHeader(http.StatusUnauthorized)
	case -3:
		// errore durante l'eliminazione
		ctx.Logger.Error("Error while deleting")
		w.WriteHeader(http.StatusUnauthorized)
	}

	path := "/photos/" + userId + "/" + photoId + ".jpg"

	// Elimino il file della foto
	err := os.Remove(path)
	if err != nil {
		ctx.Logger.Error("Error while deleting the photo")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx.Logger.Error("Photo deleted")
	w.WriteHeader(http.StatusUnauthorized)
}
