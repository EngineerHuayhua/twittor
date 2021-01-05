package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/EngineerHuayhua/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID es el id del usuario que está leyendo, page para paginar por la canitdad de resultados, search para buscar palabras, tipo es la secciond del tipo que queremos filtrar, []slice para contener n modelos de usuario a ser devueltos

//LeoUsuariosTodos lee los usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacionan conmigo
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	// siempre debe ir primero SetKip
	findOptions.SetSkip((page - 1) * 20)
	//se devolverá 20 resultados
	findOptions.SetLimit(20)

	// en el campo nombre le pasamos la condicion regex de comparacion y busqueda de strings y esta expresion (?i) indica buscara entre mayusculas y minusculas y lo adicionamos con serch que es el string de busqueda que enviamos
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	// Ahora ejecutamos el find() que devuelve la busqueda en un cursor
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	// Ahora vamos a recorrer el cursor cur y mediante Next avanzamos al siguiente registro
	for cur.Next(ctx) {
		// se debe trabajar con un modelo de usuario en particular, lo que devuelve cursor "cur" debemos decodificar en s
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		// Ahora vamos a consultar las relaciones del usuario
		var r models.Relacion
		r.UsuarioID = ID
		// ahora extraemos el string de id del cursor cur
		r.UsuarioRelacionID = s.ID.Hex()

		// por cada iteracion se decidira si incluimos o no en la respuesta
		incluir = false

		// consultamos si el usuario tiene relacion con el usurio retornado, hacemos referencia a la funcion bd/ConsultoRelacion() devuelve true o false
		encontrado, err = ConsultoRelacion(r)
		// si el tipo es new que es un busqueda nueva y no encontro entonces lo tengo que incluir porque no tiene relacion, es un usuario al cual no se sigue ******************* VERIFICAR err ****************
		if tipo == "new" && !encontrado && err != nil {
			incluir = true
		}
		// pero podria ser que sea se está siguiendo al usuario y el usuario es encontrado
		if tipo == "follow" && encontrado {
			incluir = true
		}
		// pero posiblemento me siga a mi mismo, entonces hacemos false
		if r.UsuarioRelacionID == ID {
			incluir = true
		}

		// Si incluir es true, se debe blanquear campos que no interesa incluir en el listado
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			// ahora en results agregaremos lo que tenemos almacenado en la direccion de memoria &s
			results = append(results, &s)
		}
	}

	// cuando termina el cursor preguntamos si hubo algun error interno en la operatoria del cursor, verificamos con la siguiente condicion y hubo error no devuelve nada
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	// sino hubo error cierro el cursor y devuelve el slice del modelo de usuario
	cur.Close(ctx)
	return results, true
}

// -> ./routers/listaUsuarios.go
