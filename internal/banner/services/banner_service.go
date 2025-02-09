package services

import (
	"kasikorn-line-api/internal/banner/models"
	"kasikorn-line-api/internal/banner/repositories"
)

type BannerService interface {
	GetBannerDetails(req models.BannerRequest) ([]*models.BannerResponse, error)
}

type bannerService struct {
	repo repositories.BannerRepository
}

func NewBannerService(repo repositories.BannerRepository) BannerService {
	return &bannerService{repo: repo}
}

func (s *bannerService) GetBannerDetails(req models.BannerRequest) ([]*models.BannerResponse, error) {
	banners, errResponse := s.repo.GetBannerByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	if len(banners) == 0 {
		return []*models.BannerResponse{}, nil
	}

	var bannerResponses []*models.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &models.BannerResponse{
			Title:       *banner.Title,
			Description: *banner.Description,
			Image:       *banner.Image,
		})
	}

	return bannerResponses, nil
}