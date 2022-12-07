package role

import "nurul-iman-blok-m/model"

type RoleService interface {
	SaveRole(input RoleInput) (model.Role, error)
	GetRoles(search string) ([]model.Role, error)
}

type roleService struct {
	repository RoleRepository
}

func NewRoleService(repository RoleRepository) *roleService {
	return &roleService{repository}
}

func (s *roleService) SaveRole(input RoleInput) (model.Role, error) {
	role := model.Role{}
	role.RoleName = input.RoleName
	newRole, err := s.repository.SaveRole(role)
	if err != nil {
		return role, err
	}
	return newRole, nil
}

func (s *roleService) GetRoles(search string) ([]model.Role, error) {
	if search != "" {
		roles, err := s.repository.SearchRole(search)
		if err != nil {
			return roles, err
		}
		return roles, nil
	}

	roles, err := s.repository.GetAllRole()
	if err != nil {
		return roles, err
	}

	return roles, nil
}
