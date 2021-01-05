package routers

//<- ./models/respuestaConsultaRelacion.go

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

//ConsultaRelacion verifica si hay relacion entre 2 usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	// en este EndPoint se muestra si hubo relacion true o false, a la variable modelo resp a la variable Status le asignamos false y si no fuera asi por else le asignamos true
	status, err := bd.ConsultoRelacion(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	// aquí se establece el formato json de respuesta al navegador
	w.Header().Set("Content-Type", "application/json")
	// aqui se establece el status que lleva la cabecera
	w.WriteHeader(http.StatusCreated)
	// codificamos el modelo relacion y lo enviamos al nevegador
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Error al intentar codificar el resultado de la relacion "+err.Error(), http.StatusBadRequest)
	}
}

//-> ./handlers/handlers.go
