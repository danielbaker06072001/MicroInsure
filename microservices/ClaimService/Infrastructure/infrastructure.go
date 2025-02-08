package infrastructure

import (
	model "ClaimService/Model"
	utils "ClaimService/Utils"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type ClaimRepository struct {
	db_data *gorm.DB
	db_redis *redis.Client
	db_mq *amqp.Connection
}

func NewClaimRepository(db_data *gorm.DB, db_redis *redis.Client, db_mq *amqp.Connection) *ClaimRepository {
	return &ClaimRepository{
		db_data: db_data,
		db_redis: db_redis,
		db_mq: db_mq,
	}
}

func (r *ClaimRepository) CreateClaim(claim *model.Claim) (*model.Claim, error) {
	if err := r.db_data.Model(&model.Claim{}).Create(claim).Error; err != nil {
		return nil, err
	}
	return claim, nil
}

func (r *ClaimRepository) FindAll() ([]*model.Claim, error) {
	var claims []*model.Claim
	if err := r.db_data.Find(&claims).Error; err != nil {
		return nil, err
	}
	return claims, nil
}

func (r *ClaimRepository) ValidateClaim(message string) (bool, error) {
	// Check if the claim is valid
	// * Send the claim to the Payment service to validate the claim
	queueName := "claim_validation"

	// Publish a message 
	err := utils.PublishMessage(r.db_mq, queueName, message)
	if err != nil {
		return false, err
	}
	return true, nil
}