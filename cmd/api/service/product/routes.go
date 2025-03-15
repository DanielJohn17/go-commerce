package product

import (
	"fmt"
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/types"
	"github.com/DanielJohn17/go-commerce/cmd/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/products", h.handleGetProduct)
	router.POST("/products", h.handleCreateProduct)
}

func (h *Handler) handleGetProduct(c *gin.Context) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(c *gin.Context) {
	//get JSON payload
	var newProduct types.CreateProductPayload
	if err := utils.ParseJSON(c, &newProduct); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	//validate the payload
	if err := utils.Validate.Struct(newProduct); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//create the product
	product, err := h.store.CreateProduct(types.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Image:       newProduct.Image,
		Price:       newProduct.Price,
		Quantity:    newProduct.Quantity,
	})
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusCreated, product)
}
