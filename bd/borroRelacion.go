package bd

// Para esta rutina no creamos un nuevo modelo, utilizamos el de Relacion

import (
	"context"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
)

// BorroRelacion borra la relacion en la BD
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}

//-> ./routers/bajaRelacion.go
