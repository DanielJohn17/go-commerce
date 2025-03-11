package user

import (
	"fmt"
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/service/auth"
	"github.com/DanielJohn17/go-commerce/cmd/api/types"
	"github.com/DanielJohn17/go-commerce/cmd/api/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func (h *Handler) handleRegister(c *gin.Context) {
	var newUser types.RegisterUserPayload
	if err := utils.ParseJSON(c, newUser); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmail(newUser.Email)
	if err == nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", newUser.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(c, http.StatusCreated, nil)

}
