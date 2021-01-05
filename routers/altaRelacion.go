package routers

// <- ./bd/insetoRelacion.go

import (
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

//AltaRelacion realiza el registro de la relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	// creo una variable t del tipo modelo Relacion
	var t models.Relacion
	// asignamos IDUsuario que viene del proceso token de la variable global
	t.UsuarioID = IDUsuario
	// ahora asignamos ID del usuario a quien se va a seguir que viene del parametro Request
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "NO se ha logrado insertar relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// -> ./handlers/handlers.go
