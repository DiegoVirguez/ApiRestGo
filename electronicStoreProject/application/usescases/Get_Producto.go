package usescases

import (
	"../../domain/model"
	"../../domain/ports"
	"fmt"
)

type GetProductoUseCase interface {
	Handler(codigo int64) (model.Producto, error)
}

type UseCaseGetProducto struct {
	UserRepository ports.UserRepository
}

func (useCaseGetProducto *UseCaseGetProducto) Handler(codigo int64) (model.Producto, error) {
	producto, err := useCaseGetProducto.UserRepository.GetProducto(codigo)
	if err != nil {
		fmt.Printf("server not responding %s", err.Error())
		return model.Producto{}, err
	}
	return producto, nil
}