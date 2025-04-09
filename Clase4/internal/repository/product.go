package repository

import (
	"bootcamp-web/internal/domain"
	"slices"

	"encoding/json"
	"fmt"
	"os"
)

// ProductRepository is the interface that provides product methods
type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	SearchProduct(priceGt float64) ([]domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id int, product domain.Product) (domain.Product, error)
	PatchProduct(id int, product domain.Product) (domain.Product, error)
	DeleteProduct(id int) error
}

// productRepository is a concrete implementation of ProductRepository
type productRepository struct {
	filePath string
	lastID   int
	products []domain.Product
}

// NewProductRepository creates a new ProductRepository with the necessary dependencies
func NewProductRepository(filename string) (ProductRepository, error) {
	repo := &productRepository{filePath: filename}
	err := repo.loadProducts(filename)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (pr *productRepository) loadProducts(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pr.products); err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) saveToFile() error {
	file, err := os.Create(pr.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(pr.products)
}

func (pr *productRepository) GetAllProducts() ([]domain.Product, error) {
	if pr.products == nil {
		return nil, fmt.Errorf("no products found")
	}
	return pr.products, nil
}

func (pr *productRepository) GetProductByID(id int) (domain.Product, error) {
	for _, product := range pr.products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product not found")

}

func (pr *productRepository) SearchProduct(priceGt float64) ([]domain.Product, error) {
	var filteredProducts []domain.Product
	for _, product := range pr.products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

func (pr *productRepository) CreateProduct(product domain.Product) (domain.Product, error) {
	product.ID = pr.getNextID()
	pr.products = append(pr.products, product)
	if err := pr.saveToFile(); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (pr *productRepository) getNextID() int {
	maxID := 0
	for _, product := range pr.products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) (domain.Product, error) {
	for i, p := range pr.products {
		if p.ID == id {
			product.ID = id
			pr.products[i] = product
			if err := pr.saveToFile(); err != nil {
				return domain.Product{}, err
			}
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("product not found")

}

func (pr *productRepository) PatchProduct(id int, product domain.Product) (domain.Product, error) {
	for i, p := range pr.products {
		if p.ID == id {
			if product.Name != "" {
				pr.products[i].Name = product.Name
			}
			if product.Quantity != 0 {
				pr.products[i].Quantity = product.Quantity
			}
			if product.CodeValue != "" {
				pr.products[i].CodeValue = product.CodeValue
			}
			if product.Price != 0 {
				pr.products[i].Price = product.Price
			}
			if err := pr.saveToFile(); err != nil {
				return domain.Product{}, err
			}
			return pr.products[i], nil
		}
	}

	return domain.Product{}, fmt.Errorf("product not found")

}

func (pr *productRepository) DeleteProduct(id int) error {
	for i, p := range pr.products {
		if p.ID == id {
			pr.products = slices.Delete(pr.products, i, i+1)
			// pr.products = append(pr.products[:i], pr.products[i+1:]...)
			if err := pr.saveToFile(); err != nil {
				return err
			}

			return nil
		}

	}
	return fmt.Errorf("product not found")
}
