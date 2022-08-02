package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
	"strconv"
)


type service struct{
	repository Repository
	campaignRepository campaign.Repository
	paymentService payment.Service
}


func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service)*service{
	return &service{repository, campaignRepository, paymentService}
}


type Service interface{
	GetTransactionsByCampaignID(input GetCampaignTransactionInput)([]Transaction, error)
	GetTransactionsByUserID(userID int)([]Transaction, error)
	CreateTransaction(input CreateTransactionInput)(Transaction, error)
	ProcessPayment(input TransactionNotificationInput) (error)
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

	paymentTransaction := payment.Transaction{
		ID : newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentUrl , err := s.paymentService.GetPaymentUrl(paymentTransaction,input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}

func (s *service) ProcessPayment(input TransactionNotificationInput) (error){
	transactionID, _ := strconv.Atoi(input.OrderID)
	transaction , err:=s.repository.GetByID(transactionID)
	if err != nil {
		return  err
	}

	if input.PaymentType=="credit_card" && input.TransactionStatus=="capture" && input.FraudStatus=="accept" {
		transaction.Status = "PAID"
	}else if input.TransactionStatus == "settlement"{
		transaction.Status = "PAID"
	}else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus=="cancel"{
		transaction.Status = "CANCELED"
	}

	updateTransaction, err := s.repository.Update(transaction)

	if err != nil {
		return  err
	}

	if updateTransaction.Status == "PAID"{
		campaign, err :=s.campaignRepository.FindByID(updateTransaction.CampaignID)
		if err != nil {
			return  err
		}
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updateTransaction.Amount

		_, err = s.campaignRepository.Update(campaign)
		if err != nil {
			return  err
		}
	}
	

	return nil
}
