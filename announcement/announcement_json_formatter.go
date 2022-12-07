package announcement

import "nurul-iman-blok-m/model"

type AnnouncementFormatResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Banner      string `json:"banner"`
	Slug        string `json:"slug"`
	CreatedBy   string `json:"created_by"`
}

func AnnouncementFormat(announcement model.Announcement, createdBy string) AnnouncementFormatResponse {
	return AnnouncementFormatResponse{
		ID:          announcement.ID,
		Title:       announcement.Title,
		Description: announcement.Description,
		Banner:      announcement.Images,
		Slug:        announcement.Slug,
		CreatedBy:   createdBy,
	}
}

func AnnouncementListFormat(announcement model.Announcement) AnnouncementFormatResponse {
	return AnnouncementFormatResponse{
		ID:          announcement.ID,
		Title:       announcement.Title,
		Description: announcement.Description,
		Banner:      announcement.Images,
		Slug:        announcement.Slug,
		CreatedBy:   announcement.User.Name,
	}
}

func AnnouncementsFormat(announcements []model.Announcement) []AnnouncementFormatResponse {
	formatter := []AnnouncementFormatResponse{}

	for _, announcement := range announcements {
		announcementFormatter := AnnouncementListFormat(announcement)
		formatter = append(formatter, announcementFormatter)
	}

	return formatter
}
