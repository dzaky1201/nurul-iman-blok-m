package helper

import (
	"gorm.io/gorm"
	"strconv"
)

func PaginateList(page string, perPage string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageValue, _ := strconv.Atoi(page)
		if pageValue == 0 {
			pageValue = 1
		}

		perPageValue, _ := strconv.Atoi(perPage)
		switch {
		case perPageValue > 100:
			perPageValue = 100
		case perPageValue <= 0:
			perPageValue = 10
		}

		offset := (pageValue - 1) * perPageValue // (1 - 1) * 5
		return db.Offset(offset).Limit(perPageValue)
	}
}
