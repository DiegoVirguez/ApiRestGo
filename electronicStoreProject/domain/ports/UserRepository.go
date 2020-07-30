package ports

import "../model"

type UserRepository interface {
   SaveProducto(producto *model.Producto) error
   GetProducto(codigo int64) (model.Producto, error)
   FindAllProductos()([]model.Producto, error)
   UpdateProducto(codigo int64,producto *model.Producto) (*model.Producto, error)
   DeleteProducto(codigo int64) error
}