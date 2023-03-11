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
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Prendo l'autenticazione
	auth := r.Header.Get("Authorization")

	// Prendo l'id del post a cui mettere mi piace
	phId := strings.Split(r.RequestURI, "/")[2]

	// Ottengo il testo del nuovo commento
	var comment CommentText
	err := json.NewDecoder(r.Body).Decode(&comment)

	// Controllo che il testo del commento sia valido
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, commentNotValid, http.StatusRequestEntityTooLarge)
		return
	} else if !comment.CommentTextIsValid() {
		ctx.Logger.Error(commentNotValid)
		http.Error(w, commentNotValid, http.StatusRequestEntityTooLarge)
		return
	}
	// Converto l'id utente in un int
	userId, _ := strconv.Atoi(auth)

	// Converto l'id del post in un int
	postId, _ := strconv.Atoi(phId)

	ris := rt.db.CommentPhoto(userId, postId, comment.TEXT)

	switch ris {

	case 0:
		var risultato Result
		risultato.TEXT = Done
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(risultato)

	case -1:
		ctx.Logger.Error(UserIdNotFound)
		http.Error(w, UserIdNotFound, http.StatusBadRequest)
	case -2:
		ctx.Logger.Error(photoNotFound)
		http.Error(w, photoNotFound, http.StatusProxyAuthRequired)
	case -3:
		ctx.Logger.Error(userId2Banned)
		http.Error(w, userId2Banned, http.StatusMethodNotAllowed)
	case -4:
		ctx.Logger.Error(UserIdBanned)
		http.Error(w, UserIdBanned, http.StatusForbidden)
	case -5:
		ctx.Logger.Error(ErrorServerExecution)
		http.Error(w, ErrorServerExecution, http.StatusInternalServerError)
	}
}
