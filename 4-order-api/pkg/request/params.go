package request

import (
	"fmt"
	"net/http"
	"strconv"
)

const QueryID = "id"

func GetIDFromRequest(r *http.Request) (uint, error) {
	id := r.PathValue(QueryID)
	if id == "" {
		return 0, fmt.Errorf("missing id in path: %s", r.URL.Path)
	}
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid id format: %s", id)
	}
	return uint(idUint), nil
}
