package repository

import "github.com/BatyrhanB/gin-edu/model"

type ProductsRepository interface {
	Save(products model.Product)
	Update(products model.Product)
	Delete(productsId int)
	FindById(productsId int) (products model.Product, err error)
	FindAll() []model.Product
}
