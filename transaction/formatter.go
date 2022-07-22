package transaction

import "time"

type CampaignTransactionFormatter struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Amount int `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction)CampaignTransactionFormatter{
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction)[]CampaignTransactionFormatter{
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions{
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type UserTransactionFormatter struct{
	ID int `json:"id"`
	Amount int `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct{
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction)UserTransactionFormatter{
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages)> 0{
		campaignFormatter.Name = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}


func FormatUserTransactions(transactions []Transaction)[]UserTransactionFormatter{
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionFormatter []UserTransactionFormatter
	for _, transaction := range transactions{
		formatter := FormatUserTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}