package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/announcement"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/model"
	"strconv"
	"strings"
	"time"
)

type announcementHandler struct {
	service announcement.AnnouncementService
}

func NewHandlerAnnouncement(service announcement.AnnouncementService) *announcementHandler {
	return &announcementHandler{service}
}

func (h *announcementHandler) AddAnnouncement(c *gin.Context) {
	var input announcement.AnnouncementInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("You must completed field", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	input.UserID = currentUser.ID
	input.User = currentUser

	fileImage, errBanner := c.FormFile("banner")
	if errBanner != nil {
		fmt.Println("banner-error")
		response := helper.ApiResponse("Failed to upload banner image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	extenstionFile := ""
	fileName := strings.Split(fileImage.Filename, ".")

	if len(fileName) == 2 {
		extenstionFile = fileName[1]
	}
	path := fmt.Sprintf("images/announcement-%s-%s.%s", input.Slug, time.Now().Format("2006-02-01"), extenstionFile)
	errUploadBanner := c.SaveUploadedFile(fileImage, path)

	if errUploadBanner != nil {
		fmt.Println("path-error")
		response := helper.ApiResponse("Failed to upload banner image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if currentUser.Role.RoleName == "user" || currentUser.Role.RoleName == "ustadz" {
		response := helper.ApiResponse("You not have access for add", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseAddAnnouncement, createdBy, errAdd := h.service.AddAnnouncement(input, path)
	if errAdd != nil {
		response := helper.ApiResponse("Failed to add announcement", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := announcement.AnnouncementFormat(responseAddAnnouncement, createdBy)

	response := helper.ApiResponse("Success to add announcement", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *announcementHandler) GetAllAnnouncement(c *gin.Context) {
	page := c.Request.URL.Query().Get("page")
	perPage := c.Request.URL.Query().Get("per_page")

	paginate := helper.PaginateList(page, perPage)

	announcements, count, err := h.service.GetListAnnouncement(paginate)
	if err != nil {
		response := helper.ApiResponse("Error to get announcements", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pageString, _ := strconv.Atoi(page)
	pageSizeString, _ := strconv.Atoi(perPage)

	response := helper.ApiResponseList("List Announcement", http.StatusOK, "success", pageString, pageSizeString, count, announcement.AnnouncementsFormat(announcements))
	c.JSON(http.StatusOK, response)
}

func (h *announcementHandler) GetDetailAnnouncement(c *gin.Context) {
	var input announcement.AnnouncementDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Announcement detail not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	announcementDetail, errDetail := h.service.GetDetailAnnouncement(input)
	if errDetail != nil {
		response := helper.ApiResponse("Failed to get detail announcement", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Announcement Detail", http.StatusOK, "success", announcement.AnnouncementListFormat(announcementDetail))
	c.JSON(http.StatusOK, response)
}

func (h *announcementHandler) DeleteAnnouncement(c *gin.Context) {
	var input announcement.AnnouncementDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Delete Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	if currentUser.Role.RoleName == "admin" {
		response := helper.ApiResponse("You not have access for delete", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	errDelete := h.service.DeleteAnnouncement(input)
	if errDelete != nil {
		response := helper.ApiResponse("Delete failed", http.StatusBadRequest, "error", errDelete)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Delete Success", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *announcementHandler) UpdateAnnouncement(c *gin.Context) {
	var inputID announcement.AnnouncementDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed To Update because ID not found", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUpdate announcement.AnnouncementUpdateInput
	errInputUpdate := c.ShouldBind(&inputUpdate)

	if errInputUpdate != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("You must completed field", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fileImage, _ := c.FormFile("banner")
	currentUser := c.MustGet("currentUser").(model.User)

	if fileImage != nil {
		extenstionFile := ""
		fileName := strings.Split(fileImage.Filename, ".")

		if len(fileName) == 2 {
			extenstionFile = fileName[1]
		}
		path := fmt.Sprintf("images/announcement-update-%s.%s", time.Now().Format("2006-02-01"), extenstionFile)
		errUploadBanner := c.SaveUploadedFile(fileImage, path)

		if errUploadBanner != nil {
			fmt.Println("path-error")
			response := helper.ApiResponse("Failed to upload banner image", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if currentUser.Role.RoleName == "admin" {
			response := helper.ApiResponse("You not have access for update", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		updateData, errUpdateData := h.service.UpdateAnnouncement(inputID, inputUpdate, path)
		if errUpdateData != nil {
			response := helper.ApiResponse("Failed to update announcement", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := announcement.AnnouncementFormat(updateData, updateData.User.Name)

		response := helper.ApiResponse("Success to update announcement", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)
	} else {
		if currentUser.Role.RoleName == "user" {
			response := helper.ApiResponse("You not have access for update", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		updateData, errUpdateData := h.service.UpdateAnnouncement(inputID, inputUpdate, "")
		if errUpdateData != nil {
			response := helper.ApiResponse("Failed to update announcement", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := announcement.AnnouncementFormat(updateData, updateData.User.Name)

		response := helper.ApiResponse("Success to update announcement", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)
	}
}
