package handler

import (
	"log/slog"
	"net/http"
)

func newErrorResponse(w http.ResponseWriter, logger *slog.Logger, statusCode int, message string, err error) {
	logger.Error(message)
	// @TODO: make error more pretty in postman
	http.Error(w, message+err.Error(), statusCode)
}
