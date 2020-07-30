package usescases

import (
	"../../domain/model"
	"../../domain/ports"
)

type FindAllProductosUseCase interface {
	Handler() ([]model.Producto, error)
}

type UseCaseGetFindAllProductos struct {
	UserRepository ports.UserRepository
}

func (useCaseGetFindAllProductos *UseCaseGetFindAllProductos) Handler() ([]model.Producto, error) {
	return useCaseGetFindAllProductos.UserRepository.FindAllProductos()
}