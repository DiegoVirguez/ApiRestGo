package usescases

import (
	"../commands"
	"../factory"
	"../../domain/model"
	"../../domain/ports"
)

type CreateProductoPort interface {
	Handler(userCommand commands.ProductoCommand) (model.Producto, error)
}

type UseCaseUserCreate struct {
	UserRepository ports.UserRepository
}

func (createsUseCase *UseCaseUserCreate) Handler(userCommand commands.ProductoCommand) (model.Producto, error) {

	producto, err := factory.CreateProducto(userCommand)
	if err != nil {
		return model.Producto{}, err
	}
	createProductoErr := createsUseCase.UserRepository.SaveProducto(&producto)
	if createProductoErr != nil {
		return model.Producto{}, createProductoErr
	}
	return producto, nil

}