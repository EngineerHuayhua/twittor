package routers

// <- ./bd/leoTweetsSeguidores.go

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EngineerHuayhua/twittor/bd"
)

// LeoTweetsSeguidores lee los tweets de todos nuestros seguidores
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro de página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro de página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta JSON"+err.Error(), http.StatusBadRequest)
		return
	}
}

// -> ./handlers/handlers.go
