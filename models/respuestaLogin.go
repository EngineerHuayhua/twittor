package models

// <-./jwt/jwt.go
// RespuestaLogin tiene el token que se devuelve con el login
type RespuestaLogin struct {
	// el Token empieza con mayuscula porque debe ser exportado
	// el json si debe ser en minuscula token y omitempty en caso de que exista error y si lo es, devolvemos la estructura en vacio
	Token string `json:"token,omitempty"`

	//->./handlers/handlers.go
}
