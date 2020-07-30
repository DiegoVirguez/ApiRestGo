package producto_mapper

import (
       "../../../domain/model"
       "../../adapter/repository/models"
)

func ProductoToProductoDb(producto model.Producto) models.ProductoDb {

	//now, _ := time.Parse(date.ApiDbLayout, date.GetNowDBFormatNow())
	var productoDb models.ProductoDb
	productoDb.Codigo = producto.Codigo 
	productoDb.Nombre = producto.Nombre
	productoDb.Precio = producto.Precio
	productoDb.Imagenes = producto.Imagenes
	productoDb.Caracteristicas = producto.Caracteristicas
	return productoDb
}


func ProductoDbToProducto(productoDb models.ProductoDb) model.Producto {

	//now, _ := time.Parse(date.ApiDbLayout, date.GetNowDBFormatNow())
	var producto model.Producto
	producto.Codigo  = productoDb.Codigo 
	producto.Nombre = productoDb.Nombre
	producto.Precio = productoDb.Precio
	producto.Imagenes = productoDb.Imagenes
	producto.Caracteristicas = productoDb.Caracteristicas
	return producto
}

func ProductosDbToProductos(productosDb []models.ProductoDb) []model.Producto {
	var productos []model.Producto
	for _, productoDb := range productosDb{
		producto := ProductoDbToProducto(productoDb)
		productos = append(productos, producto)
	}
	return productos
}