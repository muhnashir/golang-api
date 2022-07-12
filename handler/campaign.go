package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct{
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) (*campaignHandler){
	return &campaignHandler{service}
}


func (h *campaignHandler) GetCampaigns(c *gin.Context){
	userString := c.Query("user_id")
	userID, _ :=strconv.Atoi(userString)
	campaigns, err:=h.service.GetCampaigns(userID)

	if err != nil{
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler)GetCampaign(c *gin.Context){
	var input campaign.GetCampaignDetailInput
	
	err := c.ShouldBindUri(&input)
	
	if err!= nil{
		response := helper.APIResponse("Failed to get detail campaign",http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err!= nil{
		response := helper.APIResponse("Failed to get detail campaign",http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail",http.StatusOK,"success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) CreateCampiagn(c *gin.Context){
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	
	newCampaign, err := h.service.CreateCampiagn(input)

	if err != nil {
		response := helper.APIResponse("Create Campaign failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create Campaign Success", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler)UpdateCampign(c *gin.Context){
	var inputID campaign.GetCampaignDetailInput
	
	err := c.ShouldBindUri(&inputID)
	if err!= nil{
		response := helper.APIResponse("Failed to update campaign",http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User =  currentUser

	updateCampign, err := h.service.UpdateCampaign(inputID, inputData)
	if err!= nil{
		response := helper.APIResponse("Failed to update campaign",http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update campaign",http.StatusOK,"error", campaign.FormatCampaign(updateCampign))
	c.JSON(http.StatusOK, response)
	return
}