package app

import (
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"../../domain/ports"
	"../adapter/repository/users"
	"../database_client"
	"../controllers"
	cors "github.com/rs/cors/wrapper/gin"
	"../../application/usescases"
)

var (
	router = gin.Default()
)

func StartApplication() {

	_ = godotenv.Load()
	router.Use(cors.New(cors.Options{
	    AllowedOrigins: []string{"*"} ,
	    AllowedMethods: []string{"PUT", "GET","POST","DELETE","PATCH"},
	    AllowCredentials: true,
    }))
	userRepository := getUsersRepository()
	var handler = createHandler(userRepository)
	mapUrls(handler)
	//mapUrlLogin(createLoginHandler(userRepository))

	//logger.Info("about to start the application")
	_ = router.Run(":8081")
}

func getUsersRepository() ports.UserRepository {
	return &users.UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}

func createHandler(userRepository ports.UserRepository) controllers.RedirectProductoHandler {

	return newHandler(newCreatesUseCase(userRepository),newGetProductoUseCase(userRepository),
		   newFindAllProductosUseCase(userRepository), newUpdateProductoUseCase(userRepository),
		   newDeleteProductoUseCase(userRepository))
}

func newCreatesUseCase(repository ports.UserRepository) usescases.CreateProductoPort {
	return &usescases.UseCaseUserCreate{
		UserRepository: repository,
	}
}

func newGetProductoUseCase(repository ports.UserRepository) usescases.GetProductoUseCase {
	return &usescases.UseCaseGetProducto{
		UserRepository: repository,
	}
}

func newFindAllProductosUseCase(repository ports.UserRepository) usescases.FindAllProductosUseCase {
	return &usescases.UseCaseGetFindAllProductos{
		UserRepository: repository,
	}
}

func newUpdateProductoUseCase(repository ports.UserRepository) usescases.UpdateProductoUseCase {
	return &usescases.UseCaseUpdateProducto{
		UserRepository: repository,
	}
}

func newDeleteProductoUseCase(repository ports.UserRepository) usescases.DeleteProductoUseCase {
	return &usescases.UseCaseDeleteProducto{
		UserRepository: repository,
	}
}

func newHandler(createProducto usescases.CreateProductoPort, getProducto usescases.GetProductoUseCase,
               findAllProductos usescases.FindAllProductosUseCase, updateProducto usescases.UpdateProductoUseCase, deleteProducto usescases.DeleteProductoUseCase) controllers.RedirectProductoHandler {
	return &controllers.Handler{CreatesProductoCase: createProducto,
	                            GetProductoUseCase : getProducto,
	                            FindAllProductosUseCase : findAllProductos,
	                            UseCaseUpdateProducto : updateProducto,
	                            DeleteProductoUseCase : deleteProducto}
}

