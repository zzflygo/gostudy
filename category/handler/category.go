package handler

import (
	"category/domian/model"
	"category/domian/service"
	"category/proto/category"
	"context"
	"encoding/json"
	"go-micro.dev/v4/util/log"
)

type Category struct {
	CategoryService service.ICategoryService
}

//...创建分类
func (c *Category) CreateCategory(ctx context.Context, req *category.CategoryReq, res *category.CreateCategoryRes) error {
	cate := new(model.Category)
	err := SwapTo(req, cate)
	if err != nil {
		res.Message = "分类创建失败"
		return err
	}
	id, err := c.CategoryService.AddCategory(cate)
	if err != nil {
		res.Message = "分类创建失败"
		return err
	}
	res.Id = int64(id)
	res.Message = "分类创建成功"
	return nil
}

// ... 更新目录
func (c *Category) UpdateCategory(ctx context.Context, req *category.CategoryReq, res *category.UpdateCategoryRes) error {
	cate := new(model.Category)
	err := SwapTo(req, cate)
	if err != nil {
		res.Message = "分类更新失败"
		return err
	}
	err = c.CategoryService.UpdateCategory(cate)
	if err != nil {
		res.Message = "分类更新失败"
		return err
	}
	res.Message = "分类更新成功"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, req *category.DeleteCategoryReq, res *category.DeleteCategoryRes) error {
	err := c.CategoryService.DeleteCategory(uint64(req.Id))
	if err != nil {
		res.Message = "删除分类失败"
		return err
	}
	res.Message = "删除分类成功"
	return nil
}

func (c *Category) FindCategoryById(ctx context.Context, req *category.FindByIdReq, res *category.CategoryRes) error {
	cate, err := c.CategoryService.FindById(uint64(req.Id))
	if err != nil {
		return err
	}
	err = SwapTo(cate, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, req *category.FindByNameReq, res *category.CategoryRes) error {
	cate, err := c.CategoryService.FindByName(req.Name)
	if err != nil {
		return err
	}
	err = SwapTo(cate, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, req *category.FindAllReq, res *category.FindAllRes) error {
	cateSlice, err := c.CategoryService.FindAllCategory()
	if err != nil {
		return err
	}
	CateToSlice(cateSlice, res)
	return nil
}

func SwapTo(in, out interface{}) error {
	CategoryByte, err := json.Marshal(in)
	if err != nil {
		log.Fatal("swapTo json marshal failed:err", err)
		return err
	}
	err = json.Unmarshal(CategoryByte, out)
	if err != nil {
		log.Fatal("swapTo json unmarshal failed:err", err)
		return err
	}
	return nil
}
func CateToSlice(slice []*model.Category, res *category.FindAllRes) {
	for _, ct := range slice {
		cr := &category.CategoryRes{}
		err := SwapTo(ct, cr)
		if err != nil {
			log.Fatal("Slice swapto failed err", err)
			break
		}
		res.Category = append(res.Category, cr)
	}
	return
}
