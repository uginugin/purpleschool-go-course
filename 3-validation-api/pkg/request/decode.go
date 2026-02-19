package request

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, payload *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return nil, err
	}
	return payload, err
}
