package controllers

import (
	"github.com/fmcarrero/bookstore_oauth-go/oauth"
	"../../application/commands"
	"../../application/usescases"
	"../marshallers"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

type RedirectProductoHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	CreatesProductoCase          usescases.CreateProductoPort
	GetProductoUseCase          usescases.GetProductoUseCase
	FindAllProductosUseCase     usescases.FindAllProductosUseCase
	UseCaseUpdateProducto           usescases.UpdateProductoUseCase
	DeleteProductoUseCase       usescases.DeleteProductoUseCase
}

func (h *Handler) Create(c *gin.Context) {

	var productoCommand commands.ProductoCommand
    
	if err := c.ShouldBindJSON(&productoCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, createProductoErr := h.CreatesProductoCase.Handler(productoCommand)

	if createProductoErr != nil {
		_ = c.Error(createProductoErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, marshallers.Marshall(isPublic, result))
}

func (h *Handler) Get(c *gin.Context) {
	productoCodigo, productoErr := strconv.ParseInt(c.Param("codigo"),0,64)
	if productoErr != nil {
		restErr := rest_errors.NewBadRequestError("codigo should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	producto, errGet := h.GetProductoUseCase.Handler(productoCodigo)
	fmt.Printf("%v, %T", producto, producto)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}
    fmt.Printf("%v, %T", producto.Caracteristicas, producto.Caracteristicas)
	if oauth.GetCallerId(c.Request) == producto.Codigo {
		c.JSON(http.StatusOK, marshallers.Marshall(false, producto))
		return
	}
	c.JSON(http.StatusOK, marshallers.Marshall(oauth.IsPublic(c.Request), producto))

}

func (h *Handler) FindAll(c *gin.Context) {
	productos, err := h.FindAllProductosUseCase.Handler()
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, marshallers.MarshallArray(isPublic, productos))
}

func (h *Handler) Update(c *gin.Context) {
	codigo, codErr := strconv.ParseInt(c.Param("codigo"), 0, 64)
	if codErr != nil {
		restErr := rest_errors.NewBadRequestError("codigo should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	var productoCommand commands.ProductoCommand
	if err := c.ShouldBindJSON(&productoCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	producto, updateErr := h.UseCaseUpdateProducto.Handler(codigo, productoCommand)
	if updateErr != nil {
		restErr := rest_errors.NewBadRequestError(updateErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, &producto)
}

func (h *Handler) Delete(c *gin.Context) {
	codigo, codigoErr := strconv.ParseInt(c.Param("codigo"), 0, 64)
	if codigoErr != nil {
		restErr := rest_errors.NewBadRequestError("codigo should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	errDelete := h.DeleteProductoUseCase.Handler(codigo)
	if errDelete != nil {
		restErr := rest_errors.NewBadRequestError(errDelete.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.Status(http.StatusNoContent)
}
