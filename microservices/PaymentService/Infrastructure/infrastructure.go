package infrastructure

import (
	model "PaymentService/Model"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db_data *gorm.DB
	db_redis *redis.Client
}

func NewPaymentRepository(db_data *gorm.DB, db_redis *redis.Client) *PaymentRepository {
	return &PaymentRepository{
		db_data: db_data,
		db_redis: db_redis,
	}
}

func (r *PaymentRepository) CreatePayment(Payment *model.Payment) (*model.Payment, error) {
	if err := r.db_data.Model(&model.Payment{}).Create(Payment).Error; err != nil {
		return nil, err
	}
	return Payment, nil
}

func (r *PaymentRepository) FindAll() ([]*model.Payment, error) {
	var Payments []*model.Payment
	if err := r.db_data.Find(&Payments).Error; err != nil {
		return nil, err
	}
	return Payments, nil
}