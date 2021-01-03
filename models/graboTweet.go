package models

import "time"

//toda estructura debe empezar con un modelo
// GraboTweet es el formato o estructura que tendrá nuestro Tweet en la BD
type GraboTweet struct {
	//la minuscula userid es el nombre que tendrá en la base de datos y cuando se haga una representacion de json tendrá el nombre de userid
	UserId  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}

//->./bd/insertoTweet.go
