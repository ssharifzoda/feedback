package handlers

import (
	"feedback/internal/service"
	"feedback/pkg/logging"
	"github.com/gorilla/mux"
)

const (
	methodGet = "GET"
	methodPos = "POST"
)

type Handler struct {
	service *service.Service
	log     *logging.Logger
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *mux.Router {
	handlers := mux.NewRouter()
	handlers.HandleFunc("/countries", h.getAllCountries).Methods(methodGet)
	handlers.HandleFunc("/counlist", h.GetCountryCities).Methods(methodGet)
	handlers.HandleFunc("/serlist", h.getAllLocalServices).Methods(methodGet)
	handlers.HandleFunc("/", h.CreateFeedback).Methods(methodPos)

	return handlers
}
