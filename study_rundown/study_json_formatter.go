package study_rundown

import (
	"nurul-iman-blok-m/model"
)

type StudyRundownFormatResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	OnScheduled bool   `json:"on_scheduled"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	UstadzName  string `json:"ustadz_name"`
}

type UstadzFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func StudyResponseFormat(rundown model.StudyRundown) StudyRundownFormatResponse {

	return StudyRundownFormatResponse{
		ID:          rundown.ID,
		Title:       rundown.Title,
		OnScheduled: rundown.OnScheduled,
		Date:        rundown.ScheduleDate,
		Time:        rundown.Time,
		UstadzName:  rundown.User.Name,
	}
}

func ustadzJsonFormatter(user model.User) UstadzFormatter {
	return UstadzFormatter{
		ID:   user.ID,
		Name: user.Name,
	}
}

func ListUstadzJsonFormatter(users []model.User) []UstadzFormatter {
	listFormatter := []UstadzFormatter{}

	for _, user := range users {
		userFormatter := ustadzJsonFormatter(user)
		listFormatter = append(listFormatter, userFormatter)
	}

	return listFormatter
}

func ListRundonwnFormatter(rundowns []model.StudyRundown) []StudyRundownFormatResponse {
	formatter := []StudyRundownFormatResponse{}

	for _, rundown := range rundowns {
		rundownFormatter := StudyResponseFormat(rundown)
		formatter = append(formatter, rundownFormatter)
	}

	return formatter
}
