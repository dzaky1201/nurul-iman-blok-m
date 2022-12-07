package announcement

import "nurul-iman-blok-m/model"

type AnnouncementInput struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Slug        string `form:"slug" binding:"required"`
	UserID      uint
	User        model.User
}

type AnnouncementDetailInput struct {
	ID uint `uri:"id" binding:"required"`
}

type AnnouncementUpdateInput struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
