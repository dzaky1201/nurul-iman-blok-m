package announcement

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
	"strings"
)

type AnnouncementService interface {
	AddAnnouncement(input AnnouncementInput, imageLocation string) (model.Announcement, string, error)
	GetListAnnouncement(list func(db *gorm.DB) *gorm.DB) ([]model.Announcement, int, error)
	GetDetailAnnouncement(input AnnouncementDetailInput) (model.Announcement, error)
	DeleteAnnouncement(input AnnouncementDetailInput) error
	UpdateAnnouncement(input AnnouncementDetailInput, updateData AnnouncementUpdateInput, updatePath string, s3Client s3.Client) (model.Announcement, error)
}

type announcementService struct {
	repository AnnouncementRepository
}

func NewServiceAnnouncement(repository AnnouncementRepository) *announcementService {
	return &announcementService{repository}
}

func (s *announcementService) AddAnnouncement(input AnnouncementInput, imageLocation string) (model.Announcement, string, error) {
	announcement := model.Announcement{}
	announcement.Title = input.Title
	announcement.Description = input.Description
	announcement.Images = imageLocation
	announcement.Slug = input.Slug
	announcement.UserID = input.UserID

	announcementCreate, err := s.repository.AddAnnouncement(announcement)

	if err != nil {
		return announcementCreate, "", err
	}
	user, _ := s.repository.GetUserName(announcement, announcement.UserID)

	return announcementCreate, user.User.Name, nil
}

func (s *announcementService) GetListAnnouncement(list func(db *gorm.DB) *gorm.DB) ([]model.Announcement, int, error) {
	announcements, count, err := s.repository.GetListAnnouncement(list)
	if err != nil {
		return announcements, 0, err
	}
	return announcements, count, err
}

func (s *announcementService) GetDetailAnnouncement(input AnnouncementDetailInput) (model.Announcement, error) {
	data, err := s.repository.DetailAnnouncement(input.ID)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *announcementService) DeleteAnnouncement(input AnnouncementDetailInput) error {
	err := s.repository.DeleteAnnouncement(input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *announcementService) UpdateAnnouncement(input AnnouncementDetailInput, updateData AnnouncementUpdateInput, updatePath string, s3Client s3.Client) (model.Announcement, error) {
	data, err := s.repository.DetailAnnouncement(input.ID)
	if err != nil {
		return data, nil
	}
	if updatePath != "" {
		data.Images = updatePath
	}

	if updateData.Title != "" {
		data.Title = updateData.Title

		convertTitleToLowerCase := strings.ToLower(updateData.Title)
		sliceTitle := strings.Split(convertTitleToLowerCase, " ")
		slug := strings.Join(sliceTitle, "-")
		data.Slug = slug
	}

	if updateData.Description != "" {
		data.Description = updateData.Description
	}

	update, errUpdate := s.repository.Update(data, s3Client)
	if errUpdate != nil {
		return update, errUpdate
	}

	return update, nil
}
