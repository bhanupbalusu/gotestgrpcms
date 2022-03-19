package services

import (
	"github.com/bhanupbalusu/gotestgrpcms/domain/interface/model"
)

type ProductServicesInterface interface {
	GetProducts() (*[]model.ProductModel, error)
	GetProductByID(id string) (*model.ProductModel, error)
	CreateProduct(pm *model.ProductModel) (string, error)
	UpdateProduct(pm *model.ProductModel) error
	DeleteProduct(pm *model.ProductModel) error
}
