package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"kasikorn-line-api/internal/banner/models"
	"kasikorn-line-api/internal/banner/repositories/mock"
	repoModel "kasikorn-line-api/internal/banner/repositories/models"
	coreError "kasikorn-line-api/pkg/error"
)

func TestGetBannerDetails_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockBannerRepository(ctrl)
	bannerService := NewBannerService(mockRepo)

	req := models.BannerRequest{UserID: "user123"}
	repoBanners := []repoModel.Banner{
		{Title: strPtr("Promo 1"), Description: strPtr("Discount 10%"), Image: strPtr("image1.jpg")},
		{Title: strPtr("Promo 2"), Description: strPtr("Cashback 5%"), Image: strPtr("image2.jpg")},
	}

	mockRepo.EXPECT().GetBannerByUserID("user123").Return(repoBanners, nil)

	resp, err := bannerService.GetBannerDetails(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Promo 1", resp[0].Title)
	assert.Equal(t, "Discount 10%", resp[0].Description)
	assert.Equal(t, "image1.jpg", resp[0].Image)
}

func TestGetBannerDetails_NoBanners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockBannerRepository(ctrl)
	bannerService := NewBannerService(mockRepo)

	req := models.BannerRequest{UserID: "user123"}
	mockRepo.EXPECT().GetBannerByUserID("user123").Return([]repoModel.Banner{}, nil)

	resp, err := bannerService.GetBannerDetails(req)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	assert.Empty(t, resp)
}

func TestGetBannerDetails_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockBannerRepository(ctrl)
	bannerService := NewBannerService(mockRepo)

	req := models.BannerRequest{UserID: "user123"}

	mockError := &coreError.ErrorResponse{Message: "database error"}

	mockRepo.EXPECT().GetBannerByUserID("user123").Return(nil, mockError)

	resp, err := bannerService.GetBannerDetails(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func strPtr(s string) *string {
	return &s
}
