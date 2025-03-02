package handler

import (
	"encoding/json"
	"net/http"
	"sample-health/internal/dto"
	"sample-health/internal/model"
	"sample-health/internal/service"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		dto.WriteJSON(w, http.StatusBadRequest, response, nil)
		return
	}

	if err := h.userService.CreateUser(&user); err != nil {
		response := dto.ErrorResponse(http.StatusInternalServerError, "Failed to create user")
		dto.WriteJSON(w, http.StatusInternalServerError, response, nil)
		return
	}

	response := dto.SuccessResponse(http.StatusCreated, "User created successfully", user, nil)
	dto.WriteJSON(w, http.StatusCreated, response, nil)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		dto.WriteJSON(w, http.StatusBadRequest, response, nil)
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, "Invalid request ID")
		dto.WriteJSON(w, http.StatusBadRequest, response, nil)
		return
	}

	response := dto.SuccessResponse(http.StatusOK, "User Data.", user, nil)
	dto.WriteJSON(w, http.StatusOK, response, nil)
}
