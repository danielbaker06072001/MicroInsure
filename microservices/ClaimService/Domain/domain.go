package domain

import model "ClaimService/Model"

func NewClaim(claim *model.Claim) (*model.Claim, error) {
	return &model.Claim{
		ID:          claim.ID,
		ClaimNumber: claim.ClaimNumber,
		ClaimType:   claim.ClaimType,
		ClaimAmount: claim.ClaimAmount,
		ClaimDate:   claim.ClaimDate,
	}, nil
}

type ClaimRepository interface {
	FindAll() ([]*model.Claim, error)
	FindByID(id int) (**model.Claim, error)
	Save(claim *model.Claim) (**model.Claim, error)
}