package service

import (
	"bootcamp-web/internal/domain"
	"bootcamp-web/internal/repository"
	"errors"
	"fmt"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrInvalidInput         = errors.New("invalid input")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotCreated    = errors.New("product not created")
)

// ProductService is the interface that provides product methods
type ProductService interface {
	GetAllProducts() ([]domain.ProductDTO, error)
	GetProductByID(int) (domain.ProductDTO, error)
	SearchProduct(float64) ([]domain.ProductDTO, error)
	CreateProduct(domain.ProductRequest) (domain.ProductDTO, error)
	UpdateProduct(int, domain.ProductRequest) (domain.ProductDTO, error)
	PatchProduct(int, domain.ProductRequest) (domain.ProductDTO, error)
	DeleteProduct(int) error
}

// productService is a concrete implementation of ProductService
type productService struct {
	repository repository.ProductRepository
}

// NewProductService creates a new ProductService with the necessary dependencies
func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetAllProducts() ([]domain.ProductDTO, error) {
	product, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	if len(product) == 0 {
		return nil, ErrProductNotFound
	}
	productResponse := domain.ToDTOs(product)
	return productResponse, nil
}

func (ps *productService) GetProductByID(id int) (domain.ProductDTO, error) {
	product, err := ps.repository.GetProductByID(id)

	if err != nil {
		return domain.ProductDTO{}, ErrProductNotFound
	}
	if product.ID == 0 {
		return domain.ProductDTO{}, ErrProductNotFound
	}
	return product.ToDTO(), nil
}

func (ps *productService) SearchProduct(priceGt float64) ([]domain.ProductDTO, error) {
	var filteredProducts []domain.ProductDTO
	products, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting all  products: %w", err)
	}
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product.ToDTO())
		}
	}

	return filteredProducts, nil
}

func (ps *productService) CreateProduct(product domain.ProductRequest) (domain.ProductDTO, error) {
	if product.Name == "" || product.Quantity <= 0 || product.CodeValue == "" || product.Price <= 0 {
		return domain.ProductDTO{}, ErrInvalidInput
	}
	prod := domain.FromRequest(product)
	prod.ID = 0

	newProd, err := ps.repository.CreateProduct(prod)
	if err != nil {
		return domain.ProductDTO{}, ErrProductAlreadyExists
	}
	return newProd.ToDTO(), nil
}

func (ps *productService) UpdateProduct(id int, newProduct domain.ProductRequest) (domain.ProductDTO, error) {
	product := domain.FromRequest(newProduct)
	prod, err := ps.repository.UpdateProduct(id, product)
	if err != nil {
		return domain.ProductDTO{}, ErrProductNotFound
	}

	return prod.ToDTO(), nil

}

func (ps *productService) PatchProduct(id int, newProduct domain.ProductRequest) (domain.ProductDTO, error) {
	product := domain.FromRequest(newProduct)
	prod, err := ps.repository.PatchProduct(id, product)
	if err != nil {
		return domain.ProductDTO{}, ErrProductNotFound
	}
	return prod.ToDTO(), nil
}

func (ps *productService) DeleteProduct(id int) error {
	err := ps.repository.DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}
