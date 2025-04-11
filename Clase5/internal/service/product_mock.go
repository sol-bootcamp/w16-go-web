package service

import "bootcamp-web/internal/domain"

// mock service
type ProductServiceMock struct {
	GetAllProductsFunc func() ([]domain.ProductDTO, error)
	GetProductByIDFunc func(int) (domain.ProductDTO, error)
	SearchProductFunc  func(float64) ([]domain.ProductDTO, error)
	CreateProductFunc  func(domain.ProductRequest) (domain.ProductDTO, error)
	UpdateProductFunc  func(int, domain.ProductRequest) (domain.ProductDTO, error)
	DeleteProductFunc  func(int) error
}

func (psm *ProductServiceMock) GetAllProducts() ([]domain.ProductDTO, error) {
	if psm.GetAllProductsFunc != nil {
		return psm.GetAllProductsFunc()
	}
	return nil, nil
}

func (psm *ProductServiceMock) GetProductByID(id int) (domain.ProductDTO, error) {
	if psm.GetProductByIDFunc != nil {
		return psm.GetProductByIDFunc(id)
	}
	return domain.ProductDTO{}, nil
}

func (psm *ProductServiceMock) SearchProduct(priceGt float64) ([]domain.ProductDTO, error) {
	if psm.SearchProductFunc != nil {
		return psm.SearchProductFunc(priceGt)
	}
	return nil, nil
}

func (psm *ProductServiceMock) CreateProduct(product domain.ProductRequest) (domain.ProductDTO, error) {
	if psm.CreateProductFunc != nil {
		return psm.CreateProductFunc(product)
	}
	return domain.ProductDTO{}, nil
}

func (psm *ProductServiceMock) UpdateProduct(id int, product domain.ProductRequest) (domain.ProductDTO, error) {
	if psm.UpdateProductFunc != nil {
		return psm.UpdateProductFunc(id, product)
	}
	return domain.ProductDTO{}, nil
}

func (psm *ProductServiceMock) DeleteProduct(id int) error {
	if psm.DeleteProductFunc != nil {
		return psm.DeleteProductFunc(id)
	}
	return nil
}
