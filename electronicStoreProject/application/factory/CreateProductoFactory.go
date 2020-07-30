package factory

import (
	"../commands"
	"../../domain/model"
)

func CreateProducto(productoCommand commands.ProductoCommand) (model.Producto, error) {
	var producto model.Producto
	producto, err := producto.CreateProducto(productoCommand.Nombre, productoCommand.Precio,productoCommand.Caracteristicas)
	return producto, err
}