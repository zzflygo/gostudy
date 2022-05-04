package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"product/domian/model"
	"product/domian/service"
	"product/proto/product"
)

type Product struct {
	ProductData service.IProductService
}

func SwapTo(in, out interface{}) error {
	bytestr, err := json.Marshal(in)
	if err != nil {
		fmt.Println("SwapTo json marshal failed err:", err)
		return err
	}
	err = json.Unmarshal(bytestr, out)
	if err != nil {
		fmt.Println("SwapTo json marshal failed err:", err)
		return err
	}
	return nil
}

func (h *Product) AddProduct(ctx context.Context, request *product.ProductInfo, response *product.ProductResponse) error {
	productdata := new(model.Product)
	err := SwapTo(request, productdata)
	if err != nil {
		fmt.Println("SwapTo  failed err:", err)
		return err
	}
	id, err := h.ProductData.AddProduct(productdata)
	if err != nil {
		fmt.Println("SwapTo  failed err:", err)
		return err
	}
	response.Id = id
	response.Message = "AddProduct success..."
	return nil
}
func (h *Product) UpdateProduct(ctx context.Context, request *product.ProductInfo, response *product.ProductResponse) error {
	productdata := new(model.Product)
	err := SwapTo(request, productdata)
	if err != nil {
		fmt.Println("SwapTo  failed err:", err)
		return err
	}
	err = h.ProductData.UpdateProduct(productdata)
	if err != nil {
		fmt.Println("Update Product failed err:", err)
		return err
	}
	response.Id = productdata.ID
	response.Message = "updateproduct message success.."
	return nil
}
func (h *Product) DeleteProductById(ctx context.Context, request *product.RequestId, response *product.ProductResponse) error {
	if err := h.ProductData.DeleteProduct(request.ProductId); err != nil {
		fmt.Println("DeleteProduct failed...")
		return err
	}
	response.Id = request.ProductId
	response.Message = "DeleteProduct success...."
	return nil
}
func (h *Product) FindProductById(ctx context.Context, request *product.RequestId, response *product.ProductInfo) error {
	productdata, err := h.ProductData.FindProductByID(request.ProductId)
	if err != nil {
		fmt.Println("FindProductById failed err:", err)
		return err
	}

	if err := SwapTo(productdata, response); err != nil {
		fmt.Println("FindProductById failed err:", err)
		return err
	}
	return nil
}
func (h *Product) FindProductAll(ctx context.Context, request *product.RequestAll, response *product.AllProduct) error {
	infos := make([]*product.ProductInfo, 100)
	products, err := h.ProductData.FindProductAll()
	if err != nil {
		fmt.Println("range find all failed err:", err)
		return err
	}
	for _, v := range products {
		info := new(product.ProductInfo)
		err := SwapTo(v, info)
		if err != nil {
			fmt.Println("range find all failed err:", err)
			break
		}
		infos = append(infos, info)
	}
	response.ProductInfo = infos
	return nil
}
