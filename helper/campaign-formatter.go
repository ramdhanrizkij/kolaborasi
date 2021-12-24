package helper

import (
	"kolaborasi/entity"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign entity.Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return formatter
}

func FormatCampaigns(campaigns []entity.Campaign) []CampaignFormatter {
	formatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		formatter = append(formatter, FormatCampaign(campaign))
	}
	return formatter
}

type CampaignDetailFormatter struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	ShortDescription string          `json:"short_description"`
	ImageURL         string          `json:"image_url"`
	GoalAmount       int             `json:"goal_amount"`
	CurrentAmount    int             `json:"current_amount"`
	UserID           int             `json:"user_id"`
	User             CampaignUser    `json:"user"`
	Description      string          `json:"description"`
	Slug             string          `json:"slug"`
	Perks            []string        `json:"perks"`
	CampaignImages   []CampaignImage `json:"images"`
}

type CampaignUser struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type CampaignImage struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign entity.Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		CampaignImages:   []CampaignImage{},
	}

	if len(campaign.CampaignImages) > 0 {
		var campaignImages []CampaignImage
		formatter.ImageURL = campaign.CampaignImages[0].FileName

		for _, img := range campaign.CampaignImages {
			primary := false
			if img.IsPrimary == 1 {
				primary = true
			}
			image := CampaignImage{
				ImageURL:  img.FileName,
				IsPrimary: primary,
			}
			campaignImages = append(campaignImages, image)
		}
		formatter.CampaignImages = campaignImages
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	if campaign.User.ID != 0 {
		userCampaign := CampaignUser{
			Name:      campaign.User.Name,
			AvatarURL: campaign.User.AvatarFileName,
		}
		formatter.User = userCampaign
	}

	formatter.Perks = perks
	return formatter
}
