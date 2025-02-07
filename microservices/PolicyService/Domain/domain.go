package domain

import model "PolicyService/Model"

func NewPolicy(Policy *model.Policy) (*model.Policy, error) {
	return &model.Policy{
		ID:          Policy.ID,
		PolicyNumber: Policy.PolicyNumber,
		PolicyType:   Policy.PolicyType,
		PolicyAmount: Policy.PolicyAmount,
		PolicyDate:   Policy.PolicyDate,
	}, nil
}

type PolicyRepository interface {
	FindAll() ([]*model.Policy, error)
	FindByID(id int) (**model.Policy, error)
	Save(Policy *model.Policy) (**model.Policy, error)
}