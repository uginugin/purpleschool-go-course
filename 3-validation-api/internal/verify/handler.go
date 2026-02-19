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
	lastHash *big.Int
}

func New(router *http.ServeMux, deps *VerifyHandlerDeps) {
	handler := &verifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.send())
	router.Handle("GET /verify/{hash}", handler.verify())
}

func (h *verifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Handle[SendRequest](r)
		if err != nil {
			response.Json(w, 402, err)
			return
		}

		hash, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			response.Json(w, 402, err)
			return
		}

		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{body.email}
		e.HTML = []byte(fmt.Sprintf("<a>http://localhost:8081/verify/%d</a>", hash))
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))

		h.lastHash = hash

	}
}

func (h *verifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		isSame := hash == h.lastHash.String()

		if !isSame {
			h.lastHash = nil
		}

		response.Json(w, 200, isSame)
	}
}
