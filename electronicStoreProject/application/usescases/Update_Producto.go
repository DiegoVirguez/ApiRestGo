package usescases

import (
	"../commands"
	"../factory"
	"../../domain/model"
	"../../domain/ports"
)

type UpdateProductoUseCase interface {
	Handler(codigo int64, userCommand commands.ProductoCommand) (*model.Producto, error)
}

type UseCaseUpdateProducto struct {
	UserRepository ports.UserRepository
}

func (useCaseUpdateProducto *UseCaseUpdateProducto) Handler(codigo int64,productoCommand commands.ProductoCommand) (*model.Producto, error) {

	producto, err := factory.CreateProducto(productoCommand)
	if err != nil {
		return nil, err
	}
	productoUpd, err := useCaseUpdateProducto.UserRepository.UpdateProducto(codigo,&producto)
	if err != nil {
		return nil, err
	}
	return productoUpd, nil
}

