package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultoRelacion consulta la relacion entre 2 usuarios
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}

// para dar respuesta a la consulta como estan relacionados los usuarios se requiere crear un modelo
// -> ./models/respuestaConsultaRelacion.go
