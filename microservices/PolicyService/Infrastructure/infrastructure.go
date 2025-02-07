package infrastructure

import (
	model "PolicyService/Model"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type PolicyRepository struct {
	db_data *gorm.DB
	db_redis *redis.Client
}

func NewPolicyRepository(db_data *gorm.DB, db_redis *redis.Client) *PolicyRepository {
	return &PolicyRepository{
		db_data: db_data,
		db_redis: db_redis,
	}
}

func (r *PolicyRepository) CreatePolicy(Policy *model.Policy) (*model.Policy, error) {
	if err := r.db_data.Model(&model.Policy{}).Create(Policy).Error; err != nil {
		return nil, err
	}
	return Policy, nil
}

func (r *PolicyRepository) FindAll() ([]*model.Policy, error) {
	var Policys []*model.Policy
	if err := r.db_data.Find(&Policys).Error; err != nil {
		return nil, err
	}
	return Policys, nil
}