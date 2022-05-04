package repository

import (
	"category/domian/model"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	InitCategory() error
	CreateCategory(*model.Category) (uint64, error)
	UpdateCategory(*model.Category) error
	DeleteCategory(uint64) error
	FindCategoryById(uint64) (*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryAll() ([]*model.Category, error)
	FindCategoryByLevel(uint32) ([]*model.Category, error)
	FindCategoryByParent(string) ([]*model.Category, error)
}

type CategoryRepository struct {
	MysqlDb *gorm.DB
}

func (c *CategoryRepository) InitCategory() error {
	return c.MysqlDb.AutoMigrate(&model.Category{})

}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (c *CategoryRepository) CreateCategory(data *model.Category) (id uint64, err error) {
	err = c.MysqlDb.Model(&model.Category{}).Create(data).Error
	return data.Id, err
}
func (c *CategoryRepository) UpdateCategory(data *model.Category) error {
	return c.MysqlDb.Model(&model.Category{}).Updates(data).Error
}
func (c *CategoryRepository) DeleteCategory(id uint64) error {
	return c.MysqlDb.Model(&model.Category{}).Where("id=?", id).Delete(&model.Category{}).Error
}
func (c *CategoryRepository) FindCategoryById(id uint64) (data *model.Category, err error) {
	data = &model.Category{}
	return data, c.MysqlDb.Model(&model.Category{}).Where("id=?", id).First(data).Error
}
func (c *CategoryRepository) FindCategoryByName(name string) (data *model.Category, err error) {
	data = &model.Category{}
	return data, c.MysqlDb.Model(&model.Category{}).Where("name=?", name).First(data).Error
}

func (c *CategoryRepository) FindCategoryAll() ([]*model.Category, error) {
	data := make([]*model.Category, 128)
	return data, c.MysqlDb.Model(&model.Category{}).Find(&data).Error
}

func (c *CategoryRepository) FindCategoryByLevel(level uint32) ([]*model.Category, error) {
	data := make([]*model.Category, 128)
	return data, c.MysqlDb.Model(&model.Category{}).Where("level=?", level).Find(&data).Error
}

func (c *CategoryRepository) FindCategoryByParent(parent string) ([]*model.Category, error) {
	data := make([]*model.Category, 128)
	return data, c.MysqlDb.Model(&model.Category{}).Where("parent=?", parent).Find(&data).Error
}
