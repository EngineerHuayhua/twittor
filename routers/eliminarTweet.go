package routers

//<- ./bd/borroTweet.go

import (
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
)

// EliminarTweet permite borrar un tweet determinado
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	// pasamos id que viene de la URL y la variable global IDUsuario que viene desde el token
	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// -> ./handlers/handlers.go
