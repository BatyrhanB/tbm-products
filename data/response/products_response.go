package response

type ProductsResponse struct {
	Id           uint    `json:"id"`
	Title        string  `json:"name"`
	Description  string  `json:"description"`
	Price        int64   `json:"price"`
	InitialPrice float64 `json:"initialPrice"`
}
