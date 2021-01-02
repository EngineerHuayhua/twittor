package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/jwt"
	"github.com/EngineerHuayhua/twittor/models"
)

// <- /bd/intentoLogin.go
// Login realiza el registro
func Login(w http.ResponseWriter, r *http.Request) {
	// establecemos que el tipo de contenido que tendrá el header es json
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	//carga los datos de email y password en la varialbe t
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o Contraseña invalidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar codificar el token en un nuevo json "+err.Error(), 400)
		return
	}

	// metodo para gravar token en la cookie del usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
	// ->./jwt/jwt.go
}
