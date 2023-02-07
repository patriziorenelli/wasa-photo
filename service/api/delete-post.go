package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	// "io"
	"net/http"
	// "os"
	// "path/filepath"
	"strconv"
	"strings"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace
	userId := strings.Split(r.RequestURI, "/")[2]

	// Prendo l'id del post da eliminare
	photoId := strings.Split(r.RequestURI, "/")[4]

	if auth != userId{ 
		ctx.Logger.Error("Failed authentication")
		w.WriteHeader(http.StatusUnauthorized)
		return		
	}

	phId, _ := strconv.Atoi(photoId)
	usId, _ := strconv.Atoi(userId)

	ris := rt.db.DeletePhoto(usId, phId)

	switch ris{

		case 0:
			// tutto ok faccio eliminazione del file 

		case -1:
			// foto non esistente 
		
		case -2:
			// utente autenticato non è il proprietario del file
		


	}




}
