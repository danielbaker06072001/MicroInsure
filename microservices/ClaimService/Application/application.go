package application

import (
	infrastructure "ClaimService/Infrastructure"
	model "ClaimService/Model"
)

type ClaimService struct {
	repo infrastructure.ClaimRepository
}

func NewClaimService(repo *infrastructure.ClaimRepository) *ClaimService {
	return &ClaimService{
		repo: *repo,
	}
}

func (s *ClaimService) CreateClaim(claim *model.Claim) (*model.Claim, error) {
	claim, err := s.repo.CreateClaim(claim)
	if err != nil {
		return nil, err
	}
	return claim, nil
}

func (s *ClaimService) GetAllClaim() ([]*model.Claim, error) {
	listClaim, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return listClaim, nil
}