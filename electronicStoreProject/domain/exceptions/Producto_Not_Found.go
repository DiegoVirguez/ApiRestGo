package exceptions

type ProductoNotFound struct {
	ErrMessage string
}

func (e ProductoNotFound) Error() string {
	return e.ErrMessage
}