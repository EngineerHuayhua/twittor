package routers

//<- ./bd/leoUsuariosTodos.go

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EngineerHuayhua/twittor/bd"
)

// ListaUsuarios lee la lista de usuarios
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Cotent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Error al crear la codificacion de los resultados", http.StatusBadRequest)
		return
	}
}

//-> ./handlers/handlers.go
