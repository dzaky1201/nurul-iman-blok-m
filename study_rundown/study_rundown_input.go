package study_rundown

type StudyRundownInput struct {
	Title        string `form:"title" binding:"required"`
	OnScheduled  bool   `form:"on_scheduled" binding:"required"`
	ScheduleDate string `form:"schedule_date" binding:"required"`
	UserID       uint   `form:"user_id" binding:"required"`
	Time         string `form:"time" binding:"required"`
}

type StudyRundownInputDetail struct {
	ID uint `uri:"id" binding:"required"`
}

type StudyRundownUpdateInput struct {
	Title        string `form:"title"`
	OnScheduled  bool   `form:"on_scheduled"`
	ScheduleDate string `form:"schedule_date"`
	Time         string `form:"time"`
}
