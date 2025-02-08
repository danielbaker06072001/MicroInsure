package infrastructure

import (
	model "PaymentService/Model"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db_data *gorm.DB
	db_redis *redis.Client
	db_mq *amqp.Connection
}

func NewPaymentRepository(db_data *gorm.DB, db_redis *redis.Client, db_mq *amqp.Connection) *PaymentRepository {
	return &PaymentRepository{
		db_data: db_data,
		db_redis: db_redis,
		db_mq: db_mq,
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