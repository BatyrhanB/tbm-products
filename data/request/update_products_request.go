package request

type UpdateProductRequest struct {
	Title        string  `json:"title,omitempty" validate:"omitempty,max=155"`
	Description  string  `json:"description,omitempty" validate:"omitempty"`
	Price        int64   `json:"price,omitempty" validate:"omitempty,gt=0"` // "gt=0" ensures that price is greater than 0
	InitialPrice float64 `json:"initial_price,omitempty" validate:"omitempty,gt=0"`
}
