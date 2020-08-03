package users

import (
    "testing"
    "../../../../domain/model"
    "../models"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestCreateProducto(t *testing.T){
  var productoTest model.Producto	
  productoTest,errFactory := productoTest.CreateProducto("ejemplo1", 1552,"caracteristicas ficticias")
  
  if  errFactory != nil{
  	fmt.Println("error with connection:")
  }

  err := testEntityManager.SaveProducto(&productoTest)

  require.NoError(t,err)
  require.NotEmpty(t,productoTest.Codigo)

}

func TestReadProducto(t *testing.T){
  var productoTest =  models.ProductoDb{
  	Nombre : "ejemplo2",
  	Precio: 1600,
  	Caracteristicas : "caracteristicas ficticias",
  }	
  
  testEntityManager.Db.Create(&productoTest)

  productoGet, errorGet := testEntityManager.GetProducto(productoTest.Codigo)
	if errorGet != nil {
		fmt.Printf("server not responding %s", errorGet.Error())
	}
  
  require.NoError(t,errorGet)
  require.NotEmpty(t,productoGet.Codigo)

  require.Equal(t,productoGet.Codigo, productoTest.Codigo)
  require.Equal(t,productoGet.Nombre, productoTest.Nombre)
  require.Equal(t,productoGet.Precio, productoTest.Precio)
  require.Equal(t,productoGet.Caracteristicas, productoTest.Caracteristicas)
}

func TestFindAll(t *testing.T){
	testEntityManager.Db.Delete(&models.ProductoDb{})

	var productosTests =  []models.ProductoDb{
		models.ProductoDb{
		  	Nombre : "ejemplo1",
		  	Precio: 1600,
		  	Caracteristicas : "caracteristicas ficticias",
  		},	
  		models.ProductoDb{
		  	Nombre : "ejemplo2",
		  	Precio: 3400,
		  	Caracteristicas : "caracteristicas ficticias 2",
  		},
  		models.ProductoDb{
		  	Nombre : "ejemplo3",
		  	Precio: 2300,
		  	Caracteristicas : "caracteristicas ficticias 3",
  		},
	}

	for _, productoTest := range productosTests{
		testEntityManager.Db.Create(&productoTest)
	}

	productosResult, error := testEntityManager.FindAllProductos()

	require.NoError(t,error)
	require.Equal(t,len(productosTests),len(productosResult))
}

func TestUpdate(t *testing.T){
  var productoTest =  models.ProductoDb{
  	Nombre : "ejemplo4",
  	Precio: 5500,
  	Caracteristicas : "caracteristicas ficticias 4",
  }	
  
  testEntityManager.Db.Create(&productoTest)

  var productoUpdate =  model.Producto{
  	Codigo : productoTest.Codigo,
  	Nombre : "ejemplo_edit",
  	Precio: 4500,
  	Caracteristicas : "caracteristicas ficticias editadas",
  }

  productoResult, err := testEntityManager.UpdateProducto(productoTest.Codigo,&productoUpdate)
  
  require.NoError(t,err)
  require.NotEmpty(t,productoResult.Nombre)
  
  var productoDBResult models.ProductoDb
  testEntityManager.Db.Find(&productoDBResult, "codigo = ?", productoTest.Codigo)

  require.Equal(t,productoDBResult.Nombre,productoUpdate.Nombre)
  require.Equal(t,productoDBResult.Precio,productoUpdate.Precio)
  require.Equal(t,productoDBResult.Caracteristicas,productoUpdate.Caracteristicas)
}

func TestDelete(t *testing.T){
  var productoTest =  models.ProductoDb{
  	Nombre : "ejemplo_delete",
  	Precio: 5500,
  	Caracteristicas : "caracteristicas ficticias 4",
  }	
  
  testEntityManager.Db.Create(&productoTest)

  err := testEntityManager.DeleteProducto(productoTest.Codigo)

  require.Empty(t,err)
  
  var productoDBResult models.ProductoDb
  testEntityManager.Db.Find(&productoDBResult, "codigo = ?", productoTest.Codigo)
  
  require.Empty(t,productoDBResult)

}