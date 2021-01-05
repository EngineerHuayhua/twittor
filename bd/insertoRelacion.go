package bd

import (
	"context"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
)

// InsertoRelacion graba la relacion en la BD
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// insetamos a la BD el modelo que ya viene armado desde routers
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}

//-> ./routers/altaRelacion.go
