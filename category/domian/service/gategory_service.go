package service

import (
	"category/domian/model"
	"category/domian/repository"
)

type ICategoryService interface {
	//增
	AddCategory(*model.Category) (uint64, error)
	//删
	DeleteCategory(uint64) error
	//改
	UpdateCategory(*model.Category) error
	//查
	FindById(uint64) (*model.Category, error)
	FindByName(string) (*model.Category, error)
	FindAllCategory() ([]*model.Category, error)
	FindByParent(string) ([]*model.Category, error)
	FindByLevel(uint32) ([]*model.Category, error)
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryService(ICategoryRepository repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: ICategoryRepository,
	}
}

func (c *CategoryService) AddCategory(data *model.Category) (uint64, error) {
	return c.CategoryRepository.CreateCategory(data)
}

func (c *CategoryService) DeleteCategory(id uint64) error {
	return c.CategoryRepository.DeleteCategory(id)
}

func (c *CategoryService) UpdateCategory(data *model.Category) error {
	return c.CategoryRepository.UpdateCategory(data)
}
func (c *CategoryService) FindById(id uint64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryById(id)
}
func (c *CategoryService) FindByName(name string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(name)
}
func (c *CategoryService) FindAllCategory() ([]*model.Category, error) {
	return c.CategoryRepository.FindCategoryAll()
}
func (c *CategoryService) FindByParent(parent string) ([]*model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(parent)
}
func (c *CategoryService) FindByLevel(lv uint32) ([]*model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(lv)
}
