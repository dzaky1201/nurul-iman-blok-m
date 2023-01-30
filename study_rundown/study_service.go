package study_rundown

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type StudyService interface {
	AddStudy(input StudyRundownInput) (model.StudyRundown, error)
	GetListUstadName() ([]model.User, error)
	GetListStudy(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error)
	DetailStudy(input StudyRundownInputDetail) (model.StudyRundown, error)
	DeleteStudy(input StudyRundownInputDetail) error
	UpdateStudy(dataUpdate StudyRundownUpdateInput, input StudyRundownInputDetail) (model.StudyRundown, error)
}

type StudyServiceImpl struct {
	repository StudyRepository
}

func NewService(repository StudyRepository) *StudyServiceImpl {
	return &StudyServiceImpl{repository}
}

func (s *StudyServiceImpl) AddStudy(input StudyRundownInput) (model.StudyRundown, error) {
	study := model.StudyRundown{}
	study.Title = input.Title
	study.UserID = input.UserID
	study.Time = input.Time
	study.ScheduleDate = input.ScheduleDate
	study.OnScheduled = input.OnScheduled

	addStudy, err := s.repository.AddStudy(study)
	if err != nil {
		return model.StudyRundown{}, err
	}
	return addStudy, nil
}

func (s *StudyServiceImpl) GetListUstadName() ([]model.User, error) {
	ustadName, err := s.repository.GetListUstadName()
	if err != nil {
		return ustadName, err
	}

	return ustadName, nil
}

func (s *StudyServiceImpl) GetListStudy(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error) {
	rundowns, count, err := s.repository.GetListStudies(list)
	if err != nil {
		return rundowns, 0, err
	}
	return rundowns, count, err
}

func (s *StudyServiceImpl) DetailStudy(input StudyRundownInputDetail) (model.StudyRundown, error) {
	data, err := s.repository.DetailStudy(input.ID)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *StudyServiceImpl) DeleteStudy(input StudyRundownInputDetail) error {
	err := s.repository.DeleteStudy(input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudyServiceImpl) UpdateStudy(dataUpdate StudyRundownUpdateInput, input StudyRundownInputDetail) (model.StudyRundown, error) {
	data, err := s.repository.DetailStudy(input.ID)
	if err != nil {
		return data, nil
	}

	if dataUpdate.Title != "" {
		data.Title = dataUpdate.Title
	}
	if dataUpdate.ScheduleDate != "" {
		data.ScheduleDate = dataUpdate.ScheduleDate
	}

	if dataUpdate.Time != "" {
		data.Time = dataUpdate.Time
	}

	data.OnScheduled = dataUpdate.OnScheduled

	update, errUpdate := s.repository.UpdateStudy(data)
	if errUpdate != nil {
		return update, errUpdate
	}

	return update, nil
}
