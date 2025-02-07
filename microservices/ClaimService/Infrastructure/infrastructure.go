package infrastructure

import (
	model "ClaimService/Model"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ClaimRepository struct {
	db_data *gorm.DB
	db_redis *redis.Client
}

func NewClaimRepository(db_data *gorm.DB, db_redis *redis.Client) *ClaimRepository {
	return &ClaimRepository{
		db_data: db_data,
		db_redis: db_redis,
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