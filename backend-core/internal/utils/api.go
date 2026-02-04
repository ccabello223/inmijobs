package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data any) {
	response := dto.ResponseDTO[any]{
		Data:   data,
		Status: statusCode,
		Info:   http.StatusText(statusCode),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal error serializing response", http.StatusInternalServerError)
	}
}

func RespondError(w http.ResponseWriter, statusCode int, message string) {
	error := dto.ErrorResponseDTO{
		Message: message,
		Status:  statusCode,
		Info:    http.StatusText(statusCode),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(error); err != nil {
		http.Error(w, "Internal error serializing error response", http.StatusInternalServerError)
	}
}
