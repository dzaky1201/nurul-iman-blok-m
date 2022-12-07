package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/role"
)

type roleHandler struct {
	roleService role.RoleService
}

func NewRoleHandler(service role.RoleService) *roleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) SaveRole(c *gin.Context) {
	var input role.RoleInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Add new role failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	roleInput, errAddRole := h.roleService.SaveRole(input)
	if errAddRole != nil {
		response := helper.ApiResponse("Add new role failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := role.RoleJsonFormatter(roleInput)

	response := helper.ApiResponse("Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *roleHandler) GetRoles(c *gin.Context) {
	roleName := c.Query("role_name")

	roles, err := h.roleService.GetRoles(roleName)
	if err != nil {
		response := helper.ApiResponse("Error to get roles", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of roles", http.StatusOK, "success", role.RolesJsonFormatter(roles))
	c.JSON(http.StatusOK, response)
}
