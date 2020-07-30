package usescases

import (
	//"../../domain/model"
	"../../domain/ports"
	"fmt"
)

type DeleteProductoUseCase interface {
	Handler(codigo int64) error
}

type UseCaseDeleteProducto struct {
	UserRepository ports.UserRepository
}

func (useCaseDeleteProducto *UseCaseDeleteProducto) Handler(codigo int64) error {
	producto, err := useCaseDeleteProducto.UserRepository.GetProducto(codigo)
	if err != nil {
		fmt.Printf("server not responding %s", err.Error())
		return err
	}
	err = useCaseDeleteProducto.UserRepository.DeleteProducto(producto.Codigo)
	return err
}