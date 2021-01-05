package routers

// <- ./bd/borroRelacion.go

import (
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

// BajaRelacion realiza el borrado de la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "NO se ha logrado borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// -> ./handlers/handlers.go
