package logic

import (
	"errors"
	"time"

	"github.com/bhanupbalusu/gotestgrpcms/domain/interface/model"
	"github.com/bhanupbalusu/gotestgrpcms/domain/interface/repo"
	"github.com/bhanupbalusu/gotestgrpcms/domain/interface/services"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type productServices struct {
	ProductRepo repo.ProductRepoInterface
}

func NewProductServices(ProductRepo repo.ProductRepoInterface) services.ProductServicesInterface {
	return &productServices{ProductRepo}
}

func (p *productServices) GetProducts() (*[]model.ProductModel, error) {
	return p.ProductRepo.GetProducts()
}

func (p *productServices) GetProductByID(id string) (*model.ProductModel, error) {
	return p.ProductRepo.GetProductByID(id)
}

func (p *productServices) CreateProduct(pm *model.ProductModel) (string, error) {
	if err := validate.Validate(pm); err != nil {
		return "", errs.Wrap(ErrRedirectInvalid, "domain.application.logic.product_logic.CreateProduct")
	}
	pm.CreatedAt = time.Now().UTC().Unix()
	return p.ProductRepo.CreateProduct(pm)
}

func (p *productServices) UpdateProduct(pm *model.ProductModel) error {
	if err := validate.Validate(pm); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "domain.application.logic.product_logic.CreateProduct")
	}
	pm.UpdatedAt = time.Now().UTC().Unix()
	return p.ProductRepo.UpdateProduct(pm)
}

func (p *productServices) DeleteProduct(pm *model.ProductModel) error {
	return p.ProductRepo.UpdateProduct(pm)
}
