package repository

import (
	"errors"
	"github.com/BatyrhanB/gin-edu/data/request"
	"github.com/BatyrhanB/gin-edu/helper"
	"github.com/BatyrhanB/gin-edu/model"
	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductsRepositoryImpl(Db *gorm.DB) ProductsRepository {
	return &ProductsRepositoryImpl{Db: Db}
}

func (p *ProductsRepositoryImpl) Save(products model.Product) {
	result := p.Db.Create(&products)
	helper.ErrorPanic(result.Error)
}

func (p *ProductsRepositoryImpl) Update(products model.Product) {
	var updateProduct = request.UpdateProductRequest{
		Title:        products.Title,
		Description:  products.Description,
		Price:        products.Price,
		InitialPrice: products.InitialPrice,
	}
	result := p.Db.Model(&products).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}

func (p *ProductsRepositoryImpl) Delete(productsId int) {
	var products model.Product
	result := p.Db.Where("id = ?", productsId).Delete(&products)
	helper.ErrorPanic(result.Error)
}

func (p *ProductsRepositoryImpl) FindById(productsId int) (products model.Product, err error) {
	var product model.Product
	result := p.Db.Find(&product, productsId)
	if result.Error != nil {
		return product, nil
	} else {
		return product, errors.New("Product is not found ")
	}
}

func (p *ProductsRepositoryImpl) FindAll() []model.Product {
	var products []model.Product
	result := p.Db.Find(&products)
	helper.ErrorPanic(result.Error)
	return products
}
