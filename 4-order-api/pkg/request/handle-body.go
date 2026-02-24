package request

import "net/http"

func HandleBody[T any](r *http.Request) (*T, error) {
	var payload T

	err := Decode(r, &payload)
	if err != nil {
		return nil, err
	}

	err = Validate(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
