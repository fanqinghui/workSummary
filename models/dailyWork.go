package models

import (
	"fmt"
	_ "workSummary/utils"
)

/**
	日常工作内容
**/
type DailyWork struct {
	Id         int
	Content    string
	Year       int
	Month      int
	Day        int
	CreateDate string
}

func GetTodayWork(year, month, day int) (DailyWork, bool) {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT id,content FROM t_daliy_work where year=? and month=? and day=?", year, month, day)
	var todayWork = new(DailyWork)

	err := row.Scan(&todayWork.Id, &todayWork.Content)

	if err != nil {
		fmt.Println(err)
		return *todayWork, false
	}

	fmt.Println(todayWork.Id)

	if todayWork.Id > 0 {
		return *todayWork, true
	}

	return *todayWork, false
}
