package role

type RoleInput struct {
	RoleName string `json:"role_name" binding:"required"`
}
