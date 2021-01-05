package bd

// para borrar un documento en una coleccion no es necesario crear un modelo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BorroTweet borra un tweet determinado, ID es el id del tweet que vamos a borrar y el UserID es del usuario
func BorroTweet(ID string, UserID string) error {
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	// se debe tomar en cuenta que los campos deben ser iguales al de la base de datos, verificar la coleccion tweet con Compass
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(cxt, condicion)
	return err
}

// ->./routers/eliminarTweet.go
