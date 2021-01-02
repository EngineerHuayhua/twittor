package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/EngineerHuayhua/twittor/middlew"
	"github.com/EngineerHuayhua/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, el Handler y pongo a escuchar el Servidor */
func Manejadores() {
	// Maneja rutas que llegan por http request y realiza response
	router := mux.NewRouter()

	// EndPoint /registro que llego un tipo POST pasa el control al middLew y este revisa la BD si BD es ok devuel el control al routers.Registro
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	// <-./middlew/validoJWT.go
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(router.VerPerfil))).Methods("GET")

	// Si no existe el puerto lo crea forzandolo
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//permite realizar controles de acceso a las APIS enviado por http request
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
