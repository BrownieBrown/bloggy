package handler

import (
	"github.com/BrownieBrown/bloggy/internal/helper"
	"github.com/BrownieBrown/bloggy/internal/models"
	"net/http"
)

type HealthHandler struct {
	Config *models.Config
}

func NewHealthHandler(config *models.Config) *HealthHandler {
	return &HealthHandler{Config: config}
}

func (hh *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, "ok")
}
