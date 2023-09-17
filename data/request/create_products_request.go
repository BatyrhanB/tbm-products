package request

type CreateProductRequest struct {
	Title        string  `json:"title" validate:"required,max=155"`
	Description  string  `json:"description" validate:"required"`
	Price        int64   `json:"price" validate:"required,gt=0"` // "gt=0" ensures that price is greater than 0
	InitialPrice float64 `json:"initial_price" validate:"required,gt=0"`
}
