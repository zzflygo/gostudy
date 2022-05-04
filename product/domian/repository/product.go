package repository

import (
	"gorm.io/gorm"
	"product/domian/model"
)

type IProductRepository interface {
	CreateTable() error
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindProductAll() ([]*model.Product, error)
}

type ProductRepository struct {
	MysqlDB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}
func (p *ProductRepository) CreateTable() error {
	return p.MysqlDB.AutoMigrate(&model.Product{}, &model.ProductImage{}, &model.ProductSeo{}, &model.ProductSize{})
}
func (p *ProductRepository) AddProduct(product *model.Product) (int64, error) {
	return product.ID, p.MysqlDB.Model(&product).Create(product).Error
}
func (p *ProductRepository) DeleteProduct(id int64) error {
	tx := p.MysqlDB.Begin()
	if err := tx.Model(&model.Product{}).Unscoped().Where("id=?", id).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.ProductImage{}).Unscoped().Where("product_image_id", id).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.ProductSize{}).Unscoped().Where("product_size_id", id).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.ProductSeo{}).Unscoped().Where("product_seo_id", id).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *ProductRepository) UpdateProduct(product *model.Product) error {
	return p.MysqlDB.Model(&model.Product{}).Updates(product).Error
}
func (p *ProductRepository) FindProductByID(id int64) (product *model.Product, err error) {
	product = &model.Product{}
	err = p.MysqlDB.Model(&model.Product{}).Preload("product_image", &model.ProductImage{}).
		Preload("product_size", &model.ProductSize{}).Preload("product_seo", &model.ProductSeo{}).
		Where("id=?", id).First(product).Error
	return
}
func (p *ProductRepository) FindProductAll() (pds []*model.Product, err error) {
	pds = make([]*model.Product, 100)
	err = p.MysqlDB.Model(&model.Product{}).Preload("product_image", &model.ProductImage{}).
		Preload("product_size", &model.ProductSize{}).Preload("product_seo", &model.ProductSeo{}).Find(pds).Limit(100).Error
	return
}
