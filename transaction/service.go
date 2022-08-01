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
	GetTransactionsByUserID(userID int)([]Transaction, error)
	CreateTransaction(input CreateTransactionInput)(Transaction, error)
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

func (s *service) GetTransactionsByUserID(userID int)([]Transaction, error){
	transactions,err := s.repository.GetByCampaignByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service)CreateTransaction(input CreateTransactionInput)(Transaction, error){
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignID = input.CampaignID
	transaction.User.ID = input.User.ID
	transaction.Status = "PENDING"

	newTransaction, err := s.repository.SaveTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil

}
