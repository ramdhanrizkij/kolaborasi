package dto

import "kolaborasi/entity"

type CampaignDetailDTO struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignDTO struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             entity.User
}
