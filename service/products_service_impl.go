package service

import (
	"github.com/BatyrhanB/gin-edu/data/request"
	"github.com/BatyrhanB/gin-edu/data/response"
)

type ProductsService interface {
	Create(products request.CreateProductRequest)
	Update(products request.UpdateProductRequest)
	Delete(productsId int)
	FindById(productsId int)
	FindAll() []response.ProductsResponse
}
