package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const TOKEN = "claveSegura"

// Product representa la estructura del producto
type Product struct {
	Id    int
	Name  string
	Price float64
}

// ProductRequest representa la estructura de la solicitud para crear un producto
type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ProductResponse representa la estructura de la respuesta del producto
type ProductResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Controller contiene la "base de datos" en memoria
type productController struct {
	products map[int]Product
}

// NewController es una función constructora para Controller
func NewController() *productController {
	products := map[int]Product{
		1: {Id: 1, Name: "Notebook", Price: 1999.99},
		2: {Id: 2, Name: "Teclado", Price: 349.90},
		3: {Id: 3, Name: "Mouse", Price: 249.50},
	}
	return &productController{
		products: products,
	}
}

func main() {
	controller := NewController()
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", controller.getAllProductsHandler)
		r.Post("/", controller.createProductHandler)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func (c *productController) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Convertir el map a slice
	products := []ProductResponse{}
	for _, p := range c.products {
		productResponse := ProductResponse{
			Id:    p.Id,
			Name:  p.Name,
			Price: p.Price,
		}
		products = append(products, productResponse)
	}

	// Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (c *productController) createProductHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token != TOKEN {
		http.Error(w, "User Unauthorized", http.StatusUnauthorized)
		return
	}

	productReq := ProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// mapeo de productRequest a product
	newProduct := Product{
		Id:    len(c.products) + 1,
		Name:  productReq.Name,
		Price: productReq.Price,
	}

	// Guardar el producto en la "base de datos"
	c.products[newProduct.Id] = newProduct

	// Convertir el producto a la respuesta
	productResponse := ProductResponse{
		Id:    newProduct.Id,
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}

	// Response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productResponse)
}
