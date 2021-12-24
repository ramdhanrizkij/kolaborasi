package handler

import (
	"github.com/gin-gonic/gin"
	"kolaborasi/dto"
	"kolaborasi/entity"
	"kolaborasi/helper"
	"kolaborasi/service"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	campaignService service.CampaignService
}

func NewCampaignHandler(campaignService service.CampaignService) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaignData(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	campaign, err := h.campaignService.GetAllCampaign(userId)
	if err != nil {
		response := helper.APIResponse("error", "Failed to get campaign Data", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", "GET All Campaign Data", http.StatusOK, helper.FormatCampaigns(campaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input dto.CampaignDetailDTO

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to get campaign data", http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaign, err := h.campaignService.GetById(input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to get campaign data", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaign.ID == 0 {
		response := helper.APIResponse("error", "Campaign Not Found", http.StatusNotFound, "There is no campaign with that ID")
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.APIResponse("success", "Get Campaign Detail", http.StatusOK, helper.FormatCampaignDetail(campaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input dto.CreateCampaignDTO
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("error", "Failed to create campaign", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	input.User = currentUser
	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to create campaign", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", "Campaign Successfully Created", http.StatusOK, helper.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}
