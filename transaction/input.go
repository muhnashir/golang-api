package transaction

import "bwastartup/user"

type GetCampaignTransactionInput struct {
	ID int `uri:"id" binding:"required"`
	User       user.User
}

type CreateTransactionInput struct{
	Amount	int 	`json:"amount" binding:"required"`
	CampaignID	int	`json:"campaign_id" binding:"required"`
	User user.User
}

type TransactionNotificationInput struct{
	TransactionStatus string	`json:"transaction_status"`
	OrderID string	`json:"order_id"`
	FraudStatus	string `json:"fraud_status"`
	PaymentType string	`json:"payment_type"`
}