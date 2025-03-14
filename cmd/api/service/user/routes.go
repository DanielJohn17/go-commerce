package user

import (
	"fmt"
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/service/auth"
	"github.com/DanielJohn17/go-commerce/cmd/api/types"
	"github.com/DanielJohn17/go-commerce/cmd/api/utils"
	"github.com/DanielJohn17/go-commerce/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	// get JSON payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(c, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	}

	// find user by email
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSON(c, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(c *gin.Context) {
	// get JSON payload
	var newUser types.RegisterUserPayload
	if err := utils.ParseJSON(c, &newUser); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(newUser); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(newUser.Email)
	if err == nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", newUser.Email))
		return
	}

	// create user
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
