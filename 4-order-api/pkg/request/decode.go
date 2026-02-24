package request

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, payload *T) error {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}
	return nil
}
