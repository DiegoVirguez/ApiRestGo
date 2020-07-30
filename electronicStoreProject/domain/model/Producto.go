package model

import (
	//"github.com/fmcarrero/bookstore_users-api/domain/exceptions"
	//"github.com/fmcarrero/bookstore_users-api/domain/validators"
	//"github.com/fmcarrero/bookstore_utils-go/crypto"
	//"github.com/fmcarrero/bookstore_utils-go/date"
)

const (
	StatusActive = "active"
)

type Producto struct {
	Codigo      int64
	Nombre 	    string
	Precio      int64
	Imagenes    string
	Caracteristicas string
}

func (producto *Producto) CreateProducto(nombre string, precio int64, caracteristicas string) (Producto, error) {
	/*if err := validators.ValidateRequired(password, "Password should have some value"); err != nil {
		return User{}, exceptions.InvalidPassword{ErrMessage: err.Error()}
	}
	if err := validators.ValidateRequired(firstName, "FirstName should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateRequired(lastName, "lastName should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateRequired(email, "email should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateEmail(email, "invalid email"); err != nil {
		return User{}, exceptions.InvalidEmail{ErrMessage: err.Error()}
	}*/

	return Producto{
		Nombre:   nombre,
		Precio:    precio,
		Caracteristicas:  caracteristicas,
	}, nil
}
