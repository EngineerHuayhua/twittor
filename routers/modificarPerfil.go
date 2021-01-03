package routers

// <-./bd/modificoRegistro.go

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

// ModificarPerfil modifica el perfil de usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	// en r.Body recibimos datos json y lo decodificamos en &t
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	// a la funcion bd.ModificoRegistro le paso el modelo de usuario t codificado de body y el IDUsuario que viene capturado de middlew JWT
	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	// si status es false pero tampoco modificó ningun dato
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//->./handlers/handlers.go
