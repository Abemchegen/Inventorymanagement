package usecase

import (
	"errors"
	"inventory/domain"
)

type productUseCase struct {
	productRepository domain.ProductRepository
}

func NewProductUseCase(productRepository domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (p *productUseCase) CreateProduct(product domain.Product) (domain.Product, error) {
	if product.Name == "" {
		return domain.Product{}, errors.New("Name is required")
	}
	if product.BuyPrice == 0 {
		return domain.Product{}, errors.New("BuyPrice is required")
	}
	if product.Quantity == 0 {
		return domain.Product{}, errors.New("Quantity is required")
	}
	if product.ThresholdValue == 0 {
		return domain.Product{}, errors.New("ThresholdValue is required")
	}
	_, err := p.productRepository.CreateProduct(product)
	return product, err
}

func (p *productUseCase) GetAllProduct() ([]domain.Product, error) {
	return p.productRepository.GetAllProduct()
}

func (p *productUseCase) GetProductByID(id string) (domain.Product, error) {
	return p.productRepository.GetProductByID(id)
}

func (p *productUseCase) UpdateProduct(product domain.Product) (domain.Product, error) {
	if product.Name == "" {
		return domain.Product{}, errors.New("Name is required")
	}
	if product.BuyPrice == 0 {
		return domain.Product{}, errors.New("BuyPrice is required")
	}
	if product.Quantity == 0 {
		return domain.Product{}, errors.New("Quantity is required")
	}
	if product.ThresholdValue == 0 {
		return domain.Product{}, errors.New("ThresholdValue is required")
	}
	_, err := p.productRepository.UpdateProduct(product)
	return product, err
}

func (p *productUseCase) DeleteProduct(id string) (domain.Product, error) {
	product, err := p.productRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = p.productRepository.DeleteProduct(id)
	return product, err
}