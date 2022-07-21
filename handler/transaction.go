package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct{
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler{
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context){
	var input transaction.GetCampaignTransactionInput

	err:= c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err:=h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", transactions)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transaction", http.StatusOK, "error", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
	return
}