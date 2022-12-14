package handlers

import (
	"encoding/json"
	"feedback/internal/botSystem"
	"feedback/internal/types"
	"feedback/pkg/logging"
	"github.com/spf13/viper"
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
	log := logging.GetLogger()
	countries, err := h.service.Feedback.GetAllCountries()
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(countries)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) GetCountryCities(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()
	country := r.URL.Query().Get("countryid")
	countryId, _ := strconv.Atoi(country)
	countryCities, err := h.service.Feedback.GetCountryCities(countryId)
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(countryCities)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()
	var feedback types.Feedbacks
	var feedbackTg types.FeedBackTG
	data := r.FormValue("feedback")
	err := json.Unmarshal([]byte(data), &feedback)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 400, badRequest)
		return
	}
	reader := r.MultipartForm
	item, err := h.service.Feedback.SaveImage(reader, &feedback)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 500, err.Error())
		return
	}
	id, err := h.service.Feedback.CreateFeedback(item)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 500, internalError)
		return
	}
	feedbackTg.Massage = feedback.Massage
	feedbackTg.City = GetCitiByID(feedback.CityId)
	feedbackTg.FeedbackId = id
	feedbackTg.ChatId = viper.GetInt("chatid")
	feedbackTg.PhotoPath = feedback.Photo
	go botSystem.Sender(feedbackTg)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(thanks))
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) getAllFeedbacks(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()
	page := r.URL.Query().Get("page")
	pageNo, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 400, "invalid params")
	}
	term := r.URL.Query().Get("term")
	limit := r.URL.Query().Get("limit")
	limitNo, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 400, "invalid params")
	}
	feedbacks, err := h.service.Feedback.GetAllFeedbacks(pageNo, limitNo, term)
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(feedbacks)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) searchFeedbacks(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phone")
	page := r.URL.Query().Get("page")
	pageNo, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		NewErrorResponse(w, 400, "invalid params")
	}
	limit := r.URL.Query().Get("limit")
	limitNo, err := strconv.Atoi(limit)
	feedbacks, err := h.service.Feedback.SearchFeedbacks(phoneNumber, pageNo, limitNo)
	if err != nil {
		NewErrorResponse(w, 500, internalError)
	}
	data, err := json.Marshal(feedbacks)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}

func GetCitiByID(id int) string {
	city := map[int]string{
		1: "??????????????",
		2: "??????????????",
	}
	return city[id]
}
