package user

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type UserRepository interface {
	SaveUser(user model.User) (model.User, error)
	FindByID(ID uint) (model.User, error)
	FindByEmail(email string) (model.User, error)
	GetRoleForResponse(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) SaveUser(user model.User) (model.User, error) {
	var role model.Role
	err := r.db.Create(&user).Error
	r.db.Where("id = ?", user.RoleID).Find(&role)

	user.Role = model.Role{RoleName: role.RoleName}
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID uint) (model.User, error) {
	var user model.User
	var role model.Role
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	r.db.Where("id = ?", user.RoleID).Find(&role)
	user.Role = model.Role{RoleName: role.RoleName}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetRoleForResponse(user model.User) (model.User, error) {
	userRole := user
	err := r.db.Preload("Role").Find(&userRole).Error
	if err != nil {
		return userRole, err
	}

	return userRole, nil
}
