package handlers

import (
	"encoding/json"
	"feedback/internal/types"
	"log"
	"net/http"
	"strconv"
)

const (
	internalError = "internal error"
	badRequest    = "bad request"
	thanks        = "thanks for ur feedback"
)

func (h *Handler) getAllCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.Feedback.GetAllCountries()
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(countries)
	if err != nil {
		h.log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		h.log.Print(err)
		return
	}
}

func (h *Handler) GetCountryCities(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("countryid")
	countryId, _ := strconv.Atoi(country)
	countryCities, err := h.service.Feedback.GetCountryCities(countryId)
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(countryCities)
	if err != nil {
		h.log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		h.log.Print(err)
		return
	}
}

func (h *Handler) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback types.FeedBacks
	data := r.FormValue("feedback")
	err := json.Unmarshal([]byte(data), &feedback)
	if err != nil {
		NewErrorResponse(w, 400, badRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		NewErrorResponse(w, 400, badRequest)
		return
	}
	err = h.service.Feedback.ValidateImage(header.Size)
	if err != nil {
		NewErrorResponse(w, 400, badRequest)
		return
	}
	filename := header.Filename
	item, err := h.service.Feedback.SaveImage(file, filename, &feedback)
	if err != nil {
		NewErrorResponse(w, 500, err.Error())
		return
	}
	err = h.service.Feedback.CreateFeedback(item)
	if err != nil {
		NewErrorResponse(w, 500, internalError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(thanks))
	if err != nil {
		log.Print(err)
		return
	}
}
