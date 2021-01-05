package bd

// <- ./models/devuelvoTweetsSeguidores.go

import (
	"context"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// LeoTweetsSeguidores lee los tweets de mis seguidores
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	// para verificar con que usarios está relacionado
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	// creamos un slice de tipo bson para almacenar muchas condiciones
	condiciones := make([]bson.M, 0)
	//match buscara el usuario Id de la relacion dentro de la relacion, le pasamos otra condicion usuarioid tiene que ser igual a ID
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	// con el resultado de la anaterior condicion vamos a unir con los tweets con la instruccion loopup y los 4 parametros para unir tablas de mongo, from = con que tabla queremos unir la tabla relacion, atravez de campo se une usuariorelacionid, en la coleccion tweet como se llama ese campo "foreignField": "userid"
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	// nos permite que todos los documentos vengan exactamente iguales, el usuario con su tweet
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	// los datos que vengan serán ordenados $sort por tweet.fecha del mas actual al mas antiguo -1 y 1 lo contrario
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	// cuantos registro tiene que saltar en la busqueda
	condiciones = append(condiciones, bson.M{"$skip": skip})
	// cuantos registros se debe leer para ser mostrado
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var result []models.DevuelvoTweetsSeguidores

	// se ejecutar directamente en la base de datos con las condiciones dadas y se creara un cursor que viene con todos los datos y no es necesario recorrerlo con un for
	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return result, false
	}
	// procesamos armado de todos los registros que tenemos que enviar en &result y exise un error se devolverá en err
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}

// -> ./routers/leoTweetsSeguidores.go
