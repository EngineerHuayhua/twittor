package bd

import (
	"context"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya está en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	// en resultado se guarda el modelo de usuario y esta trae ID que es de usuario y lo estamos convertiendo el ID en string
	ID := resultado.ID.Hex()
	if err != nil {
		// No encontró la condicion especificada
		return resultado, false, ID
	}
	// devuelve que si existe el usuario
	return resultado, true, ID
}
