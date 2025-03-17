package cart

import (
	"fmt"
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/service/auth"
	"github.com/DanielJohn17/go-commerce/cmd/api/types"
	"github.com/DanielJohn17/go-commerce/cmd/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(
	store types.OrderStore,
	productStore types.ProductStore,
	userStore types.UserStore) *Handler {
	return &Handler{store, productStore, userStore}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/cart/checkout", auth.WithJWTAuth(h.HandleCheckout, h.userStore))
}

func (h *Handler) HandleCheckout(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c.Request.Context())
	// parse JSON payload
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(c, &cart); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	//validate the payload
	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// get products
	productIDs, err := getCartItemIDs(cart.Items)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	products, err := h.productStore.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
