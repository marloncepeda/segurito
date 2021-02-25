package quotationservice

import (
    "src/models/quotation"
	"src/repositories/quotationrepo"
)

type QuotationService interface {
	CreateQuotation(quotation *quotation.Quotation) (*quotation.Quotation, error)
}

type quotationService struct {
	Repo quotationrepo.Repo
}

// NewQuotationService will instantiate User Service
func NewQuotationService(
	repo quotationrepo.Repo,
) QuotationService {

	return &quotationService{
		Repo: repo,
	}
}

func (qs *quotationService) CreateQuotation(quotation *quotation.Quotation) (*quotation.Quotation, error) {
	return qs.Repo.CreateQuotation(quotation)
}
