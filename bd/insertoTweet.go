package bd

// <-./models/graboTweet.go

import (
	"context"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoTweet graba el Tweet en al BD
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserId,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	// en result se almacena tipo json o bson resultado de insertar un registro
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	// del campo result.InsertedID extrae la clave del ultimo campo insertado y obtiene el ObjectID
	objID, _ := result.InsertedID.(primitive.ObjectID)
	// existe dos maneras de hacerlo uno es con String() y otro Hex()
	return objID.String(), true, nil
}
