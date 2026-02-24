package product

import (
	"4-order-api/pkg/handler"
	"4-order-api/pkg/request"
	"4-order-api/pkg/response"
	"net/http"
)

type Handler struct {
	repo *ProductRepo
}

type HandlerDeps struct {
	Repo *ProductRepo
}

func NewHandler(r *http.ServeMux, deps *HandlerDeps) {
	h := &Handler{repo: deps.Repo}

	handler.HandleCRUD(r, "/products", h)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.repo.Get(id)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	response.Encode(w, http.StatusOK, product)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "failed to get products", http.StatusInternalServerError)
		return
	}

	response.Encode(w, http.StatusOK, products)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	product, err := request.HandleBody[Product](r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.repo.Create(product)
	if err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	response.Encode(w, http.StatusCreated, product)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.repo.Get(id)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	updatedProduct, err := request.HandleBody[Product](r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.repo.Update(id, updatedProduct)
	if err != nil {
		http.Error(w, "failed to update product", http.StatusInternalServerError)
		return
	}

	response.Encode(w, http.StatusOK, updatedProduct)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
