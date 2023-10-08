package endpoints

import (
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/errors"

	"context"
	"encoding/json"
	"net/http"
)

func defaultEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if e, ok := response.(errors.Errorer); ok && e.Error() != nil {
		encodeError(ctx, e.Error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	statusCode := ErrorStatusCode(err)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func ErrorStatusCode(err error) int {
	var statusCode int

	switch err {
	case errors.ErrMissingId:
		statusCode = http.StatusBadRequest
	case errors.ErrBadIdFormat:
		statusCode = http.StatusBadRequest
	case errors.ErrBadRosterFormat:
		statusCode = http.StatusBadRequest
	case errors.ErrIdNotFound:
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}
