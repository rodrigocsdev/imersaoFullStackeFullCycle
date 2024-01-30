package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodrigocsdev/imersaoFullStackeFullCycle/goapi/internal/entity"
	"github.com/rodrigocsdev/imersaoFullStackeFullCycle/goapi/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ProductService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: ProductService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "category_id")
	if categoryID == "" {
		http.Error(w, "category_ID is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
