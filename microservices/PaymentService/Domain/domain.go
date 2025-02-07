package domain

import model "PaymentService/Model"

func NewPayment(Payment *model.Payment) (*model.Payment, error) {
	return &model.Payment{
		ID:          Payment.ID,
		PaymentNumber: Payment.PaymentNumber,
		PaymentType:   Payment.PaymentType,
		PaymentAmount: Payment.PaymentAmount,
		PaymentDate:   Payment.PaymentDate,
	}, nil
}

type PaymentRepository interface {
	FindAll() ([]*model.Payment, error)
	FindByID(id int) (**model.Payment, error)
	Save(Payment *model.Payment) (**model.Payment, error)
}