package usecase

import (
	"errors"
	"inventory/domain"
)

type StoreUseCase struct {
	StoreRepository domain.StoreRepository
}

func NewStoreUseCase(StoreRepository domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{
		StoreRepository: StoreRepository,
	}
}

func (s *StoreUseCase) CreateStore(store domain.Store) (domain.Store, error) {
	if store.Name == "" {
		return domain.Store{}, errors.New("name is required")
	}

	return s.StoreRepository.CreateStore(store)
}

func (s *StoreUseCase) GetAllStore() ([]domain.Store, error) {
	return s.StoreRepository.GetAllStore()
}

func (s *StoreUseCase) GetStoreByID(id string) (domain.Store, error) {
	return s.StoreRepository.GetStoreByID(id)
}

func (s *StoreUseCase) UpdateStore(store domain.Store) (domain.Store, error) {
	if store.Name == "" {
		return domain.Store{}, errors.New("name is required")
	}

	return s.StoreRepository.UpdateStore(store)
}

func (s *StoreUseCase) DeleteStore(id string) (domain.Store, error) {
	return s.StoreRepository.DeleteStore(id)
}


