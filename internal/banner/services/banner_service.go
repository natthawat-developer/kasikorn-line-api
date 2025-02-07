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
	// ดึงข้อมูล Banner โดยใช้ userID
	banners, errResponse := s.repo.GetBannerByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	// สร้าง Response ที่จะส่งกลับในรูปแบบ List ของ BannerResponse
	var bannerResponses []*models.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &models.BannerResponse{
			BannerID:   banner.BannerID,
			Title:      *banner.Title,
			Description: *banner.Description,
			Image:      *banner.Image,
		})
	}

	return bannerResponses, nil
}
