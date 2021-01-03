package bd

//<-./models/devuelvoTweets.go

import (
	"context"
	"log"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeoTweets lee tweets de un perfil
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	bd := MongoCN.Database("twittor")
	col := bd.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	// se busca en la BD la coleccion tweet, tenemos un campo llamado userid que tiene que coincidir con el ID que le vamos a pasar
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	// para listar tweets de 20
	opciones.SetLimit(20)
	// ahora se debe ordenar por campo fecha y -1 de forma descendente del documento a la hora de ordenar
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	// para saltar pagina de 20
	opciones.SetSkip((pagina - 1) * 20)

	//creamos un puntero (como la tabal de BD) donde se van a guardar los resultados de a 1 por cada registro encontrado se tiene que armar el resultado
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		// muestra error por consola
		log.Fatal(err.Error())
		return resultados, false
	}

	// avanzar al siguiente registro del contexto el TODO crea un contexto vacio sin ningun tipo de limitacion
	for cursor.Next(context.TODO()) {
		// para trabajar cada tweet en particular
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}

//-> ./routers/leoTweets.go
