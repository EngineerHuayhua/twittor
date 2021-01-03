package routers

//<-./models/tweet.go

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

// GraboTweet permite grabar el Tweet en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	// creamos una variable mensaje para almacenar el mensaje que viene desde el body
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error al decodificar el mensaje tweet "+err.Error(), 400)
		return
	}

	// armamos un registro con el IDUsuario que viene del token, asignamos el mensaje recibido y registramos la hora actual
	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	// el registro lo pasamo de parametro a al rutina del paquete bd que devuelve status y error
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}

	// si no logro a grabar nada
	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//->./handlers/handlers.go
