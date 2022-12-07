package role

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type RoleRepository interface {
	SaveRole(input model.Role) (model.Role, error)
	GetAllRole() ([]model.Role, error)
	SearchRole(search string) ([]model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) SaveRole(role model.Role) (model.Role, error) {
	err := r.db.Save(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) GetAllRole() ([]model.Role, error) {
	var role []model.Role
	err := r.db.Find(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) SearchRole(search string) ([]model.Role, error) {
	var role []model.Role
	err := r.db.Where("role_name like ?", "%"+search+"%").Find(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}
