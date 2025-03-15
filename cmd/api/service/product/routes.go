package product

import (
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/types"
	"github.com/DanielJohn17/go-commerce/cmd/api/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/products", h.handleGetProduct)
}

func (h *Handler) handleGetProduct(c *gin.Context) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, products)
}
