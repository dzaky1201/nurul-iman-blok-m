package study_rundown

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type StudyRepository interface {
	AddStudy(study model.StudyRundown) (model.StudyRundown, error)
	GetListUstadName() ([]model.User, error)
	GetListStudies(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error)
	DetailStudy(ID uint) (model.StudyRundown, error)
	DeleteStudy(ID uint) error
	UpdateStudy(study model.StudyRundown) (model.StudyRundown, error)
}

type StudyRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *StudyRepositoryImpl {
	return &StudyRepositoryImpl{db}
}

func (s *StudyRepositoryImpl) AddStudy(study model.StudyRundown) (model.StudyRundown, error) {
	var user model.User
	err := s.db.Create(&study).Error

	s.db.Where("id = ?", study.UserID).Find(&user)
	study.User = model.User{Name: user.Name}

	if err != nil {
		return study, err
	}
	return study, nil
}

func (s *StudyRepositoryImpl) GetListUstadName() ([]model.User, error) {
	var users []model.User
	var role model.Role
	var ustadzName []model.User

	err := s.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	for _, user := range users {
		s.db.Where("id = ?", user.RoleID).Find(&role)
		itemUstadzname := model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  model.Role{RoleName: role.RoleName},
		}
		if itemUstadzname.Role.RoleName == "ustadz" {
			ustadzName = append(ustadzName, itemUstadzname)
		}
		role = model.Role{}
	}

	return ustadzName, nil

}

func (s *StudyRepositoryImpl) GetListStudies(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error) {
	var rundowns []model.StudyRundown
	var user model.User
	var listsStudyRundowns []model.StudyRundown

	err := s.db.Scopes(list).Find(&rundowns).Error

	for _, item := range rundowns {
		s.db.Where("id = ?", item.UserID).Find(&user)
		itemRundown := model.StudyRundown{
			ID:           item.ID,
			Title:        item.Title,
			OnScheduled:  item.OnScheduled,
			ScheduleDate: item.ScheduleDate,
			User:         model.User{Name: user.Name},
			UserID:       item.UserID,
			Time:         item.Time,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}

		listsStudyRundowns = append(listsStudyRundowns, itemRundown)
		user = model.User{}
	}

	if err != nil {
		return listsStudyRundowns, 0, err
	}
	totalCount := int64(0)
	s.db.Find(&rundowns).Count(&totalCount)
	return listsStudyRundowns, int(totalCount), nil
}

func (s *StudyRepositoryImpl) DeleteStudy(ID uint) error {
	var studyRundown model.StudyRundown

	s.db.Where("id = ?", ID).Find(&studyRundown)
	err := s.db.Delete(&model.StudyRundown{}, ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *StudyRepositoryImpl) DetailStudy(ID uint) (model.StudyRundown, error) {
	var studyRundown model.StudyRundown
	err := s.db.Preload("User").Where("id = ?", ID).Find(&studyRundown).Error
	if err != nil {
		return studyRundown, err
	}
	return studyRundown, nil
}

func (s *StudyRepositoryImpl) UpdateStudy(study model.StudyRundown) (model.StudyRundown, error) {
	err := s.db.Save(&study).Error
	if err != nil {
		return study, err
	}
	return study, nil
}
