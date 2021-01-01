package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCD es el objeto de conexion a la BD y se puede realizar todas las operaciones con la base de datos*/
var MongoCD = ConectarBD()
var clientOption = options.Client().ApplyURI("mongodb+srv://dbDonatelo:Dnt81H10h29@cluster0.tv0gk.mongodb.net/twittor?retryWrites=true&w=majority")

/* ConectarBD es la funcion que me permite conectar a la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/* ChequeoConnection es un ping a la BD */
func ChequeoConnection() int {
	err := MongoCD.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
