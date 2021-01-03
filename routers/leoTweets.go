package routers

//<- ./bd/leoTweets.go

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EngineerHuayhua/twittor/bd"
)

//LeoTweets lee los Tweets
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id ", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página ", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	//convertimos pagina de tipo int a int64 debido a que la rutina ./bd/leoTweets requiere este tipo de dato
	pag := int64(pagina)
	// ./bd/leoTweets() devuelve un modelo "devuelvoTweets" y un boleano
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los Tweets", http.StatusBadRequest)
		return
	}
	// setting de tipo json
	w.Header().Set("Content-Type", "application/json")
	// e indicamos que fue satisfactorio
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, "Hubo un error al momento de codificar en formato json", http.StatusBadRequest)
		return
	}
}

// -> ./handlers/handlers.go
