package models

//import "time"

type ProductoDb struct {
	Codigo      int64 `gorm:"primary_key"`
	Nombre 	    string
	Precio      int64
	Imagenes    string
	Caracteristicas    string
}

func (ProductoDb) TableName() string {
	return "productos"
}