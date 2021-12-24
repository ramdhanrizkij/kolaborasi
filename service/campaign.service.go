package service

import (
	"fmt"
	"github.com/gosimple/slug"
	"kolaborasi/dto"
	"kolaborasi/entity"
	"kolaborasi/repository"
)

type CampaignService interface {
	GetAllCampaign(userID int) ([]entity.Campaign, error)
	GetById(input dto.CampaignDetailDTO) (entity.Campaign, error)
	CreateCampaign(input dto.CreateCampaignDTO) (entity.Campaign, error)
}

type campaignService struct {
	campaignRepo repository.CampaignRepository
}

func NewCampaignService(repo repository.CampaignRepository) CampaignService {
	return &campaignService{repo}
}

func (s *campaignService) GetAllCampaign(userID int) ([]entity.Campaign, error) {

	if userID == 0 {
		campaign, err := s.campaignRepo.FindAllCampaign()
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}

	campaign, err := s.campaignRepo.FindByUserID(userID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *campaignService) GetById(input dto.CampaignDetailDTO) (entity.Campaign, error) {
	campaign, err := s.campaignRepo.FindById(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *campaignService) CreateCampaign(input dto.CreateCampaignDTO) (entity.Campaign, error) {
	campaign := entity.Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserID:           input.User.ID,
	}

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	slugCampaign := slug.Make(slugCandidate)
	campaign.Slug = slugCampaign

	newCampaign, err := s.campaignRepo.SaveCampaign(campaign)
	if err != nil {
		return campaign, err
	}
	return newCampaign, nil

}
