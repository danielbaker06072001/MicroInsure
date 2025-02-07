package application

import (
	infrastructure "PolicyService/Infrastructure"
	model "PolicyService/Model"
)

type PolicyService struct {
	repo infrastructure.PolicyRepository
}

func NewPolicyService(repo *infrastructure.PolicyRepository) *PolicyService {
	return &PolicyService{
		repo: *repo,
	}
}

func (s *PolicyService) CreatePolicy(Policy *model.Policy) (*model.Policy, error) {
	Policy, err := s.repo.CreatePolicy(Policy)
	if err != nil {
		return nil, err
	}
	return Policy, nil
}

func (s *PolicyService) GetAllPolicy() ([]*model.Policy, error) {
	listPolicy, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return listPolicy, nil
}