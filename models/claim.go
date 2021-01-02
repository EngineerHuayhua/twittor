package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// <-./handlers/handlers.go
// Claim es la estructura usada para procesar el JWT que viene para desemcriptarlo
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}

//->./middlew/validoJWT.go
