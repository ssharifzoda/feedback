package handlers

import (
	"feedback/pkg/logging"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, massage string) {
	logger := logging.GetLogger()
	logger.Println(massage)
	http.Error(w, massage, statusCode)
}
