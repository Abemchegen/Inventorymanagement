package usecase

import (
	"inventory/domain"
)

type SupplierUseCase struct {
	supplierRepository domain.SuppliersRepository
}

func NewSupplierUseCase(SupplierRepository domain.SuppliersRepository) domain.SuppliersUseCase {
	return &SupplierUseCase{
		supplierRepository: SupplierRepository,
	}


}

func (s *SupplierUseCase) CreateSuppliers(suppliers domain.Suppliers) (domain.Suppliers, error) {

	result, err := s.supplierRepository.CreateSuppliers(suppliers)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return result, nil
}

func (s *SupplierUseCase) GetAllSuppliers() ([]domain.Suppliers, error) {
	result, err := s.supplierRepository.GetAllSuppliers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SupplierUseCase) GetSuppliersByID(id string) (domain.Suppliers, error) {
	result, err := s.supplierRepository.GetSuppliersByID(id)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return result, nil
}

func (s *SupplierUseCase) DeleteSuppliers(id string) (domain.Suppliers, error) {
	result, err := s.supplierRepository.DeleteSuppliers(id)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return result, nil
}

func (s *SupplierUseCase) UpdateSuppliers(suppliers domain.Suppliers) (domain.Suppliers, error) {
	result, err := s.supplierRepository.UpdateSuppliers(suppliers)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return result, nil
}


