package handler

import (
	"bootcamp-web/internal/domain"
	"bootcamp-web/internal/service"
	"bootcamp-web/pkg/apperrors"
	"bootcamp-web/pkg/web"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := ph.service.GetAllProducts()
		if err != nil {
			web.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(w, http.StatusOK, "products found", products)

	}
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	product, err := ph.service.GetProductByID(id)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	web.Success(w, http.StatusOK, "product found", product)
}

func (ph *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	if priceGtStr == "" {
		web.Error(w, http.StatusBadRequest, "priceGt is required")
		return
	}

	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	products, err := ph.service.SearchProduct(priceGt)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	web.Success(w, http.StatusOK, "products found", products)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		web.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := ph.service.CreateProduct(newProduct)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	web.Success(w, http.StatusCreated, "product created", product)

}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var newProduct domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		web.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validar el producto
	if newProduct.CodeValue == "" || newProduct.Name == "" || newProduct.Price <= 0 {
		web.Error(w, http.StatusBadRequest, "invalid product data")
		return
	}

	// Enviar el producto a actualizar al servicio
	product, err := ph.service.UpdateProduct(id, newProduct)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, apperrors.ErrInternalServer.Error())
		return
	}

	web.Success(w, http.StatusOK, "product updated", product)

}

func (ph *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var newProduct domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		web.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Enviar el producto a actualizar al servicio
	product, err := ph.service.PatchProduct(id, newProduct)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, apperrors.ErrInternalServer.Error())
		return
	}

	web.Success(w, http.StatusOK, "product updated partially", product)

}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = ph.service.DeleteProduct(id)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, apperrors.ErrNotFound.Message)
		return
	}

	web.Success(w, http.StatusNoContent, "product deleted", nil)

}
