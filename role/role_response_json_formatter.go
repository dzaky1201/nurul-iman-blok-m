package role

import "nurul-iman-blok-m/model"

type RoleFormatter struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

func RoleJsonFormatter(role model.Role) RoleFormatter {
	return RoleFormatter{
		ID:       role.ID,
		RoleName: role.RoleName,
	}
}

func RolesJsonFormatter(roles []model.Role) []RoleFormatter {
	rolesFormatter := []RoleFormatter{}

	for _, role := range roles {
		roleFormatter := RoleJsonFormatter(role)
		rolesFormatter = append(rolesFormatter, roleFormatter)
	}

	return rolesFormatter
}
