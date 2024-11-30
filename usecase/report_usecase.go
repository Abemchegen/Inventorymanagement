package usecase

import (
	"inventory/domain"
)

type ReportUseCase struct {
	productRepository domain.ReportRepository
}

func NewReportUseCase(reportRepository domain.ReportRepository) domain.ReportUseCase {
	return &ReportUseCase{
		productRepository: reportRepository,
	}
}

func (a *ReportUseCase) GetBestSellingProduct() (domain.Report, error) {
	reports, err := a.productRepository.GetBestSellingProduct()
	if err != nil {
		return domain.Report{}, err
	}
	return reports, nil
}

func (a *ReportUseCase) GetBestSellingCategory() ([]domain.BestSellingCategory, error) {
	catagories, err := a.productRepository.GetBestSellingCategory()
	if err != nil {
		return nil, err
	}
	return catagories, nil
}

func (a *ReportUseCase) GetOverView() (domain.OverView, error) {
	overview, err := a.productRepository.GetOverView()
	if err != nil {
		return domain.OverView{}, err
	}
	return overview, nil
}