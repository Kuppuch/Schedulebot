package database

import (
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	gorm.Model
	Name   string
	Start  time.Time
	End    time.Time
	Source string
	Type   string
	Day    string
}

func GetCurrentLessons() []Lesson {
	weekday, weektype := getCurrentDay()
	var lessons []Lesson
	//DB.Where("(type = ? OR type = ?) AND day = ? AND CAST(start AS time) between CAST(now() AS time) AND addtime(now(), \"0:15:00\")", weektype, "всегда", weekday).Find(&lessons)
	DB.Where("(type = ? OR type = ?) AND day = ? AND start::time between (now()::time + interval '3 hour') AND (now()::time + interval '15 minutes' + interval '3 hour')", weektype, "всегда", weekday).Find(&lessons)
	return lessons
}

func GetToday() []Lesson {
	weekday, weektype := getCurrentDay()
	var lessons []Lesson
	DB.Where("(type = ? OR type = ?) AND day = ?", weektype, "всегда", weekday).Find(&lessons)
	return lessons
}

func getCurrentDay() (string, string) {
	now := time.Now()
	_, week := now.ISOWeek()

	weekday := now.Weekday().String()
	switch weekday {
	case "Monday":
		weekday = "понедельник"
	case "Tuesday":
		weekday = "вторник"
	case "Wednesday":
		weekday = "среда"
	case "Thursday":
		weekday = "четверг"
	case "Friday":
		weekday = "пятница"
	case "Saturday":
		weekday = "суббота"
	case "Sunday":
		weekday = "воскресенье"
	}

	weektype := ""
	switch week % 2 {
	case 0:
		weektype = "знаменатель"
	case 1:
		weektype = "числитель"
	}

	return weekday, weektype
}
