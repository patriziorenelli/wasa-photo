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
		return
	} else if !comment.CommentTextIsValid() {

		w.WriteHeader(http.StatusBadRequest)
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
		risultato.TEXT = "Done"
		risultato.CODE = 200
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(risultato)
		return

	case -1:
		ctx.Logger.Error("User not exist")
		w.WriteHeader(http.StatusUnauthorized)
		return

	case -2:
		ctx.Logger.Error("The post doesn't exist")
		w.WriteHeader(http.StatusUnauthorized)
		return

	case -3:
		ctx.Logger.Error("You banned other user")
		w.WriteHeader(http.StatusUnauthorized)
		return

	case -4:
		ctx.Logger.Error("The other user blocked you")
		w.WriteHeader(http.StatusUnauthorized)
		return

	case -5:
		ctx.Logger.Error("Error during execution")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

}
