package users

import (
	"errors"
	"fmt"
	"../../../../domain/exceptions"
	"../../../../domain/model"
	"../models"
	"../../../mappers/producto_mapper"
	//"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserMysqlRepository struct {
	Db *gorm.DB
}

func (userMysqlRepository *UserMysqlRepository) SaveProducto(producto *model.Producto) error {

	var productoDb models.ProductoDb
	productoDb = producto_mapper.ProductoToProductoDb(*producto)
	if err := userMysqlRepository.Db.Create(&productoDb).Error; err != nil {
		//logger.Error(fmt.Sprintf("can't work with %s", userDb.FirstName), err)
		return errors.New(fmt.Sprintf("can't work with %s", productoDb.Nombre))
	}
	producto.Codigo = productoDb.Codigo
	//user.Password = ""
	return nil
}

func (userMysqlRepository *UserMysqlRepository) GetProducto(codigo int64) (model.Producto, error) {
	var productoDb models.ProductoDb
	if userMysqlRepository.Db.Find(&productoDb, "codigo = ?", codigo).Error != nil {
		return model.Producto{}, exceptions.ProductoNotFound{ErrMessage: fmt.Sprintf("product with codigo=%d not found", codigo)}
	}
	producto := producto_mapper.ProductoDbToProducto(productoDb)
	return producto, nil
}

func (userMysqlRepository *UserMysqlRepository) FindAllProductos()([]model.Producto, error) {
	var productosDb []models.ProductoDb
	if userMysqlRepository.Db.Limit(30).Find(&productosDb).Error != nil {
		return nil, errors.New(fmt.Sprintf("error en la consulta"))
	}
	if len(productosDb) <= 0 {
		return nil, errors.New(fmt.Sprintf("no hay usuarios registrados"))
	}
	productos := producto_mapper.ProductosDbToProductos(productosDb)
	return productos, nil
}

func (userMysqlRepository *UserMysqlRepository) UpdateProducto(codigo int64,producto *model.Producto) (*model.Producto, error){

	var productoUpd models.ProductoDb
	if userMysqlRepository.Db.First(&productoUpd, codigo).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("user not found %v", codigo))
	}
	if userMysqlRepository.Db.Model(&productoUpd).Where("codigo = ?", codigo).Update(models.ProductoDb{Nombre : producto.Nombre, Precio: producto.Precio, Caracteristicas: producto.Caracteristicas}).Error != nil {
		return nil, errors.New(fmt.Sprintf("error when updated producto %v", codigo))
	}
	productoUpdated := producto_mapper.ProductoDbToProducto(productoUpd)
	return &productoUpdated, nil
}

func (userMysqlRepository *UserMysqlRepository) DeleteProducto(codigo int64) error{
   var producto models.ProductoDb
	producto.Codigo = codigo
	if userMysqlRepository.Db.Delete(&producto).Error != nil {
		return errors.New(fmt.Sprintf("cannot delete producto  %v", codigo))
	}
	return nil
}
