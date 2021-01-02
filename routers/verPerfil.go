package routers

// <-./bd/buscoPerfil.go

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
)

// VerPerfil permite extraer los valores del Perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(perfil)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar codificar el perfil "+err.Error(), 400)
		return
	}
}
