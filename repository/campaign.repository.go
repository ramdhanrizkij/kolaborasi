package repository

import (
	"gorm.io/gorm"
	"kolaborasi/entity"
)

type CampaignRepository interface {
	FindAllCampaign() ([]entity.Campaign, error)
	FindByUserID(userID int) ([]entity.Campaign, error)
	FindById(ID int) (entity.Campaign, error)
	SaveCampaign(campaign entity.Campaign) (entity.Campaign, error)
}
type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) FindAllCampaign() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns)
	if err != nil {
		return campaigns, err.Error
	}
	return campaigns, nil
}

func (r *campaignRepository) FindByUserID(userID int) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := r.db.Where("user_id=?", userID).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns)
	if err != nil {
		return campaigns, err.Error
	}
	return campaigns, nil
}

func (r *campaignRepository) FindById(ID int) (entity.Campaign, error) {
	var campaign entity.Campaign

	err := r.db.Preload("CampaignImages").Preload("User").Find(&campaign, ID)
	if err != nil {
		return campaign, err.Error
	}
	return campaign, nil
}

func (r *campaignRepository) SaveCampaign(campaign entity.Campaign) (entity.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
