package models

// Relacion modelo para grabar la relacion de un usuario con otro
type Relacion struct {
	// UsuarioID es el Id de usuario del que está logeado
	UsuarioID string `bson:"usuarioid" json:"usuarioId"`
	// UsuarioRelacionID son los Id de otros usuarios con las cuales se tiene relacion
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}

//-> ./bd/insertoRelacion.go
