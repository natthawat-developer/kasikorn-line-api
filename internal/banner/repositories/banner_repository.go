package repositories

import (
	"errors"
	"kasikorn-line-api/internal/banner/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type BannerRepository interface {
	GetBannerByUserID(userID string) ([]models.Banner, *error.ErrorResponse)
}

type bannerRepository struct {
	DB *gorm.DB
}

func NewBannerRepository(DB *gorm.DB) BannerRepository {
	return &bannerRepository{DB: DB}
}


func (r *bannerRepository) GetBannerByUserID(userID string) ([]models.Banner, *error.ErrorResponse) {
	var banners []models.Banner
	err := r.DB.Where(&models.Banner{UserID: &userID}).Find(&banners).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return banners, nil
}
