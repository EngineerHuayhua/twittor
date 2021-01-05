package routers

// Previamente se debería tener creado en la ruta ./db/modificoRegistro.go como ya lo tenemos no es necesario volver a crearlo

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
)

// SubirBanner sube el Banner al servidor
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	// lo trataremos como un formulario de html y vendra en un parametro (archivo) llamado banner
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error al recibir el archivo banner "+err.Error(), http.StatusBadRequest)
		return
	}
	// el archivo lo separamos en dos a partir del punto y devolvemos un string solo la seccion de extension
	var extension = strings.Split(handler.Filename, ".")[1]
	// aqui cresmos una variable archivo con el nombre y la extension
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	// abrimos y creamos el archivo con los permisos 0666 en el disco duro
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	// copiamos el archivo que llegó del Request file y lo alojamos en f
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

//-> ./handlers/handlers.go
