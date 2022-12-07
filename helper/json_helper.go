package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Info Info        `json:"info"`
	Data interface{} `json:"data"`
}

type ResponseList struct {
	Info InfoList    `json:"info"`
	Data interface{} `json:"data"`
}

type Info struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type InfoList struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Count   int    `json:"count"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	info := Info{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Info: info,
		Data: data,
	}

	return jsonResponse
}

func ApiResponseList(message string, code int, status string, page int, pageSize int, count int, data interface{}) ResponseList {
	info := InfoList{
		Message: message,
		Code:    code,
		Status:  status,
		Page:    page,
		PerPage: pageSize,
		Count:   count,
	}

	jsonResponse := ResponseList{
		Info: info,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
