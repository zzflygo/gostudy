package service

import (
	"product/domian/model"
	"product/domian/repository"
)

type IProductService interface {
	AddProduct(*model.Product) (int64, error)
	UpdateProduct(*model.Product) error
	DeleteProduct(int64) error
	FindProductByID(int64) (*model.Product, error)
	FindProductAll() ([]*model.Product, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (p *ProductService) AddProduct(product *model.Product) (int64, error) {
	return p.ProductRepository.AddProduct(product)
}

func (p *ProductService) UpdateProduct(product *model.Product) error {
	return p.ProductRepository.UpdateProduct(product)
}

func (p *ProductService) DeleteProduct(id int64) error {
	return p.ProductRepository.DeleteProduct(id)
}
func (p *ProductService) FindProductByID(id int64) (product *model.Product, err error) {
	product = &model.Product{}
	product, err = p.ProductRepository.FindProductByID(id)
	return
}
func (p *ProductService) FindProductAll() (pdts []*model.Product, err error) {
	pdts = make([]*model.Product, 100)
	pdts, err = p.ProductRepository.FindProductAll()
	return
}
