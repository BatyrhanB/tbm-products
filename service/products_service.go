package service

import (
	"github.com/BatyrhanB/gin-edu/data/request"
	"github.com/BatyrhanB/gin-edu/data/response"
	"github.com/BatyrhanB/gin-edu/helper"
	"github.com/BatyrhanB/gin-edu/model"
	"github.com/BatyrhanB/gin-edu/repository"
	"github.com/go-playground/validator/v10"
)

type ProductsServiceImpl struct {
	ProductRepository repository.ProductsRepository
	Validate          *validator.Validate
}

func (p *ProductsServiceImpl) Create(products request.CreateProductRequest) {
	err := p.Validate.Struct(products)
	helper.ErrorPanic(err)
	productModel := model.Product{
		Title: products.Title,
	}
	p.ProductRepository.Save(productModel)
}

func (p *ProductsServiceImpl) Update(products request.UpdateProductRequest) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductsServiceImpl) Delete(productsId int) {
	p.ProductRepository.Delete(productsId)
}

func (p *ProductsServiceImpl) FindById(productsId int) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductsServiceImpl) FindAll() []response.ProductsResponse {
	result := p.ProductRepository.FindAll()

	var products []response.ProductsResponse
	for _, value := range result {
		product := response.ProductsResponse{
			Id:           value.ID,
			Title:        value.Title,
			Description:  value.Description,
			Price:        value.Price,
			InitialPrice: value.InitialPrice,
		}
		products = append(products, product)
	}
	return products
}

func NewProductsServiceImpl(productsRepository repository.ProductsRepository, validate *validator.Validate) ProductsService {
	return &ProductsServiceImpl{
		ProductRepository: productsRepository,
		Validate:          validate,
	}
}
