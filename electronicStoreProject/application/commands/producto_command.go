package commands

type ProductoCommand struct {
	Nombre     string `json:"nombre"`
	Precio     int64 `json:"precio"`
	Caracteristicas  string `json:"caracteristicas"`
	Imagenes  string `json:"imagenes"`
}