package service

import (
	"github.com/farinchan/thesis-attendance-backend/model"
	"github.com/farinchan/thesis-attendance-backend/repository"
)

func GetStudentAttendanceRules() ([]model.StudentAttendanceRule, error) {
	return repository.GetAllStudentAttendanceRules()
}

func GetStudentAttendanceRule(id uint) (model.StudentAttendanceRule, error) {
	return repository.GetStudentAttendanceRuleByID(id)
}

func GetStudentAttendanceRulesByDay(day string) (model.StudentAttendanceRule, error) {
	return repository.GetStudentAttendanceRuleByDay(day)
}
