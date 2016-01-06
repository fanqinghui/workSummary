package models

import (
	"fmt"
	"workSummary/utils"
)

/**
	日常工作内容
**/
type DailyWork struct {
	Id         int
	UId        int
	Content    string
	Year       int
	Month      int
	Day        int
	CreateDate string
}

func GetTodayWork(year, month, day int) (DailyWork, bool) {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT id,content FROM t_daily_work where year=? and month=? and day=?", year, month, day)
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

/**
保存今日工作日报
**/
func SaveTodayWork(work DailyWork) ReturnJson {
	db := OpenDB()
	defer db.Close()
	var returnJson ReturnJson

	if work.Id > 0 { //update
		stmt, err := db.Prepare("update t_daily_work set content=? where id=?")
		defer stmt.Close()
		utils.CheckErr(err)
		res, err := stmt.Exec(work.Content, work.Id)
		utils.CheckErr(err)
		line, _ := res.RowsAffected()
		fmt.Println("更新影响行数", line)

		returnJson = ReturnJson{
			Id:      int64(work.Id),
			Message: "更新成功",
		}
	} else { //insert
		stmt, err := db.Prepare("insert into t_daily_work(uid,content,year,month,day) values(?,?,?,?,?)")
		defer stmt.Close()
		utils.CheckErr(err)
		res, err := stmt.Exec(work.UId, work.Content, work.Year, work.Month, work.Day)
		utils.CheckErr(err)
		id, err := res.LastInsertId()
		utils.CheckErr(err)

		returnJson = ReturnJson{
			Id:      int64(id),
			Message: "保存成功",
		}
	}
	return returnJson
}
