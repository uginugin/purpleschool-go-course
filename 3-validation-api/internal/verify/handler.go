package verify

import (
	"3-validation-api/config"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/response"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	*config.Config
}

type verifyHandler struct {
	*config.Config
	HashMail map[string]string
}

func New(router *http.ServeMux, deps *VerifyHandlerDeps) {
	handler := &verifyHandler{
		Config:   deps.Config,
		HashMail: make(map[string]string),
	}
	router.HandleFunc("POST /send", handler.send())
	router.Handle("GET /verify/{hash}", handler.verify())
}

func (h *verifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Handle[SendRequest](r)
		if err != nil {
			response.Json(w, 402, err.Error())
			return
		}

		hash, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			response.Json(w, 402, err.Error())
			return
		}

		e := email.NewEmail()
		e.From = fmt.Sprintf("Jordan Wright <%s>", h.Config.Email)
		e.To = []string{body.Email}
		e.HTML = []byte(fmt.Sprintf("<a href=\"http://localhost:8081/verify/%d\">Подтвердить email</a>", hash))
		err = e.Send(
			fmt.Sprintf("%s:587", h.Config.Address),
			smtp.PlainAuth("", h.Config.Email, h.Config.Password, h.Config.Address))

		if err != nil {
			response.Json(w, 500, err.Error())
			return
		}

		h.HashMail[hash.String()] = body.Email
		response.Json(w, 200, hash.String())

	}
}

func (h *verifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		response.Json(w, 200, h.HashMail[hash] != "")
		delete(h.HashMail, hash)
	}
}
