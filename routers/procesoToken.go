package routers

import (
	"errors"
	"strings"

	"github.com/EngineerHuayhua/twittor/bd"
	"github.com/EngineerHuayhua/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//<-./handlers/handlers.go

// Email valor de Email usado en todos lo EndPoints
var Email string

// IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints
var IDUsuario string

// ProcesoToken procesa token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	// Split divide el texto a partir de tk delimitado por Bearer y esta se vuelve en vector de 2 indices
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	// TrimSpace quita los espacios que puede tener el tk
	tk = strings.TrimSpace(splitToken[1])

	// mapea dentro claims y verifica que el token sea valido
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	// si no hubo error, el tkn es valido se verifica que el email es valido en la BD
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Inválido")
	}

	return claims, false, string(""), err
}

// ->./bd/buscoPerfil.go
