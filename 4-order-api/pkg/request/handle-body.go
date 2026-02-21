package request

import "net/http"

func HandleBody[T any](r *http.Request, w *http.ResponseWriter) (*T, error) {
	var payload T

	body, err := Decode(r, &payload)
	if err != nil {
		return nil, err
	}

	err = Validate(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
