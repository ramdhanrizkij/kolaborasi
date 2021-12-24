package service

import (
	"kolaborasi/dto"
	"kolaborasi/entity"
	"kolaborasi/repository"
)

type CampaignService interface {
	GetAllCampaign(userID int) ([]entity.Campaign, error)
	GetById(input dto.CampaignDetailDTO) (entity.Campaign, error)
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
