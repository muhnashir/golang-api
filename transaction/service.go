package transaction

import (
	"bwastartup/campaign"
	"errors"
)



type service struct{
	repository Repository
	campaignRepository campaign.Repository
}


func NewService(repository Repository, campaignRepository campaign.Repository)*service{
	return &service{repository, campaignRepository}
}


type Service interface{
	GetTransactionsByCampaignID(input GetCampaignTransactionInput)([]Transaction, error)
}

func(s *service) GetTransactionsByCampaignID(input GetCampaignTransactionInput)([]Transaction, error){
	campaign, err:= s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	
	transaction, err:= s.repository.GetByCampaignByID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
