package handler

import (
	"4-order-api/pkg/request"
	"fmt"
	"net/http"
)

type CRUDHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func HandleCRUD(r *http.ServeMux, basePath string, h CRUDHandler) {
	r.HandleFunc(fmt.Sprintf("%s/%s", basePath, request.QueryID), h.Get)
	r.HandleFunc(fmt.Sprintf("%s/", basePath), h.GetAll)
	r.HandleFunc(fmt.Sprintf("POST %s/", basePath), h.Create)
	r.HandleFunc(fmt.Sprintf("PATCH %s/%s", basePath, request.QueryID), h.Update)
	r.HandleFunc(fmt.Sprintf("DELETE %s/%s", basePath, request.QueryID), h.Delete)
}
