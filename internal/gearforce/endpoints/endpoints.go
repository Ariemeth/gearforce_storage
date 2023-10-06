package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
)

func defaultEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
