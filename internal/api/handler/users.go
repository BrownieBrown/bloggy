package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/BrownieBrown/bloggy/internal/database"
	"github.com/BrownieBrown/bloggy/internal/helper"
	"github.com/BrownieBrown/bloggy/internal/models"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	Config *models.Config
}

func NewUserHandler(cfg *models.Config) *UserHandler {
	return &UserHandler{
		Config: cfg,
	}
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		helper.RespondWithError(w, http.StatusBadRequest, "request body is empty")
		return
	}

	defer r.Body.Close()

	var params CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      params.Name,
	}

	newUser, err := uh.Config.ApiConfig.DB.CreateUser(r.Context(), createUserParams)
	if err != nil {
		log.Printf("Error creating user: %v", err) // Log the error
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := UserResponse{
		ID:        newUser.ID,
		CreatedAt: newUser.CreatedAt.Time,
		UpdatedAt: newUser.UpdatedAt.Time,
		Name:      newUser.Name,
	}

	helper.RespondWithJSON(w, http.StatusCreated, resp)
}
