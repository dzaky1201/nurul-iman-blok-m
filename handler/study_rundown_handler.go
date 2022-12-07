package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/model"
	"nurul-iman-blok-m/study_rundown"
	"strconv"
)

type StudyRundownHandler struct {
	service study_rundown.StudyService
}

func NewHandlerStudyRundown(service study_rundown.StudyService) *StudyRundownHandler {
	return &StudyRundownHandler{service}
}

func (h *StudyRundownHandler) AddStudy(c *gin.Context) {
	var input study_rundown.StudyRundownInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("you must complete field", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)

	if currentUser.Role.RoleName == "admin" {
		response := helper.ApiResponse("You not have access for add", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	study, errAdd := h.service.AddStudy(input)
	if errAdd != nil {
		response := helper.ApiResponse("Failed to add rundown", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := study_rundown.StudyResponseFormat(study)

	response := helper.ApiResponse("Success to add rundown", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) GetListUstadzName(c *gin.Context) {
	name, err := h.service.GetListUstadName()
	if err != nil {
		response := helper.ApiResponse("Error to get ustadz name", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List ustadz", http.StatusOK, "success", study_rundown.ListUstadzJsonFormatter(name))
	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) GetAllRundown(c *gin.Context) {
	page := c.Request.URL.Query().Get("page")
	perPage := c.Request.URL.Query().Get("per_page")

	paginate := helper.PaginateList(page, perPage)

	listStudy, count, err := h.service.GetListStudy(paginate)
	if err != nil {
		response := helper.ApiResponse("Error to get rundown", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pageString, _ := strconv.Atoi(page)
	pageSizeString, _ := strconv.Atoi(perPage)

	response := helper.ApiResponseList("List Rundown", http.StatusOK, "success", pageString, pageSizeString, count, study_rundown.ListRundonwnFormatter(listStudy))
	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) GetDetailStudyRundown(c *gin.Context) {
	var input study_rundown.StudyRundownInputDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Rundown detail not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	studyRundown, errDetail := h.service.DetailStudy(input)
	if errDetail != nil {
		response := helper.ApiResponse("Failed to get detail Rundown", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Rundown Detail", http.StatusOK, "success", study_rundown.StudyResponseFormat(studyRundown))
	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) DeleteStudyRundown(c *gin.Context) {
	var input study_rundown.StudyRundownInputDetail
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
	errDelete := h.service.DeleteStudy(input)
	if errDelete != nil {
		response := helper.ApiResponse("Delete failed", http.StatusBadRequest, "error", errDelete)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Delete Success", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) UpdateStudyRundown(c *gin.Context) {
	var inputID study_rundown.StudyRundownInputDetail
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed To Update because ID not found", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUpdate study_rundown.StudyRundownUpdateInput
	errInputUpdate := c.ShouldBind(&inputUpdate)

	if errInputUpdate != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("You must completed field", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)

	if currentUser.Role.RoleName == "user" {
		response := helper.ApiResponse("You not have access for update", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updateData, errUpdateData := h.service.UpdateStudy(inputUpdate, inputID)
	if errUpdateData != nil {
		response := helper.ApiResponse("Failed to update announcement", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := study_rundown.StudyResponseFormat(updateData)

	response := helper.ApiResponse("Success to update study rundown", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
