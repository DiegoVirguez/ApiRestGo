package marshallers

import (
	"encoding/json"
	"fmt"
	"../../domain/model"
)

type PublicProducto struct {
	Codigo          int64  `json:"codigo"`
	Nombre string `json:"nombre"`
	Precio      int64 `json:"precio"`
}
type PrivateProducto struct {
	Codigo          int64  `json:"codigo"`
	Nombre string `json:"nombre"`
	Precio      int64 `json:"precio"`
	Caracteristicas string `json:"caracteristicas"`
}

func Marshall(isPublic bool, producto model.Producto) interface{} {
	if isPublic {
		return PublicProducto{
			Codigo:          producto.Codigo,
			Nombre: producto.Nombre,
			Precio:      producto.Precio,
		}
	}
	productoJson, errUn := json.Marshal(producto)
	fmt.Println(errUn)
	fmt.Println(producto)
	var privateProducto PrivateProducto
	_ = json.Unmarshal(productoJson, &privateProducto)
	return privateProducto
}

func MarshallArray(isPublic bool, productos []model.Producto) []interface{} {
	result := make([]interface{}, len(productos))
	for index, producto := range productos {
		result[index] = Marshall(isPublic, producto)
	}
	return result
}