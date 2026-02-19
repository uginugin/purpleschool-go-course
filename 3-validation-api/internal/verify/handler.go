package verify

import (
	"3-validation-api/config"
	"net/http"
)

type VerifyHandlerDeps struct {
	*config.Config
}

type verifyHandler struct {
	*config.Config
}

func NewVerifyHandler(router *http.ServeMux, deps *VerifyHandlerDeps) {
	handler := &verifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.send())
	router.Handle("GET /verify/{hash}", handler.verify())
}

func (h *verifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *verifyHandler) verify() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
