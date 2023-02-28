package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var risultato Result

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
		risultato.TEXT = Done

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

	if ris != 0 {
		return
	}

	mydir, err := os.Getwd()
	if err != nil {
		ctx.Logger.Error("Error during directory creation")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	path := mydir + "/photos/" + userId + "/" + photoId + ".jpg"
	// Elimino il file della foto
	err = os.Remove(path)
	if err != nil {
		ctx.Logger.Error("Error while deleting the photo")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(risultato)
}
