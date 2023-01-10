package api

import (
	// "encoding/json"
	"fmt"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// DA FARE
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// LIMITE A 10MB
	r.ParseMultipartForm(10 * 1024 * 1024)
	fmt.Print(r)
	file, handler, err := r.FormFile("myFile")

	fmt.Print("AIUTO")
	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()

	fmt.Println("File info")
	fmt.Print("MADONNA")
	fmt.Println("File name: ", handler.Filename)
	fmt.Println("File size: ", handler.Size)
	fmt.Println("File type: ", handler.Header.Get("Content-Type"))

}
