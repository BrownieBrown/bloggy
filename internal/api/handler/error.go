package handler

import (
	"github.com/BrownieBrown/bloggy/internal/helper"
	"github.com/BrownieBrown/bloggy/internal/models"
	"net/http"
)

type ErrorHandler struct {
	Config *models.Config
}

func NewErrorHandler(config *models.Config) *ErrorHandler {
	return &ErrorHandler{Config: config}
}

func (eh *ErrorHandler) HandleError(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithError(w, http.StatusInternalServerError, "error")
}
