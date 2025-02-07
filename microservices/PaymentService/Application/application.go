package application

import (
	infrastructure "PaymentService/Infrastructure"
	model "PaymentService/Model"
)

type PaymentService struct {
	repo infrastructure.PaymentRepository
}

func NewPaymentService(repo *infrastructure.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: *repo,
	}
}

func (s *PaymentService) CreatePayment(Payment *model.Payment) (*model.Payment, error) {
	Payment, err := s.repo.CreatePayment(Payment)
	if err != nil {
		return nil, err
	}
	return Payment, nil
}

func (s *PaymentService) GetAllPayment() ([]*model.Payment, error) {
	listPayment, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return listPayment, nil
}