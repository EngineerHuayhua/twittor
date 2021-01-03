package jwt

import (
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// <- ./routers/login.go
// GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	// creamos una nueva clave privada
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	// establecemos los privilegios y jamas se debe guardar el password del modelo de usuario en jwt
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	// creamos un nuevo objeto de jwt le pasamos el header SigningMethodHS256 y payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// ahora el objeto creado token se debe firmar
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

	// ->./models/respuestaLogin.go
}
