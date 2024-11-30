package usecase

import (
	"errors"
	"inventory/domain"
)

type OrdersUseCaseUseCase struct {
	productRepository domain.OrdersRepository
}

func NewOrderUseCase(OrderRepository domain.OrdersRepository) domain.OrdersUseCase {
	return &OrdersUseCaseUseCase{
		productRepository: OrderRepository,
	}
}


func (a *OrdersUseCaseUseCase) CreateOrders(orders domain.Orders) (domain.Orders, error) {
	if orders.ProductName == "" {
		return domain.Orders{}, errors.New("oroduct name is empty")
	}
	if orders.OrderValue == 0 {
		return domain.Orders{}, errors.New("order value is empty")
	}
	if orders.Quantity == 0 {
		return domain.Orders{}, errors.New("quantity is empty")
	}
	result, err := a.productRepository.CreateOrders(orders)
	if err != nil {
		return domain.Orders{}, err
	}
	return result, nil
}

func (a *OrdersUseCaseUseCase) GetAllOrders() ([]domain.Orders, error) {
	result, err := a.productRepository.GetAllOrders()
	if err != nil {
		return []domain.Orders{}, err
	}
	return result, nil
}

func (a *OrdersUseCaseUseCase) UpdateOrders(orders domain.Orders) (domain.Orders, error) {
	if orders.ProductName == "" {
		return domain.Orders{}, errors.New("product name is empty")
	}
	if orders.OrderValue == 0 {
		return domain.Orders{}, errors.New("order value is empty")
	}
	if orders.Quantity == 0 {
		return domain.Orders{}, errors.New("quantity is empty")
	}
	result, err := a.productRepository.UpdateOrders(orders)
	if err != nil {
		return domain.Orders{}, err
	}
	return result, nil
}

func (a *OrdersUseCaseUseCase) DeleteOrders(id string) (domain.Orders, error) {
	result, err := a.productRepository.DeleteOrders(id)
	if err != nil {
		return domain.Orders{}, err
	}
	return result, nil
}

func (a *OrdersUseCaseUseCase) GetOrdersByID(id string) (domain.Orders, error) {
	result, err := a.productRepository.GetOrdersByID(id)
	if err != nil {
		return domain.Orders{}, err
	}
	return result, nil
}
