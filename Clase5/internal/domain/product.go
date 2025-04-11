package domain

type ProductDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

// ToDTO converts a Product to a ProductDTO
func (p *Product) ToDTO() ProductDTO {
	return ProductDTO{
		ID:          p.ID,
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  p.Expiration,
		Price:       p.Price,
	}
}

// ToDTOs converts a slice of Product to a slice of ProductDTO
func ToDTOs(products []Product) []ProductDTO {
	dtos := make([]ProductDTO, len(products))
	for i, p := range products {
		dtos[i] = p.ToDTO()
	}
	return dtos
}

// FromRequest creates a Product from a ProductRequest
func FromRequest(req ProductRequest) Product {
	return Product{
		Name:        req.Name,
		Quantity:    req.Quantity,
		CodeValue:   req.CodeValue,
		IsPublished: req.IsPublished,
		Expiration:  req.Expiration,
		Price:       req.Price,
	}
}
