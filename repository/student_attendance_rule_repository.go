package repository

import (
	"github.com/farinchan/thesis-attendance-backend/config"
	"github.com/farinchan/thesis-attendance-backend/model"
)

func GetAllStudentAttendanceRules() ([]model.StudentAttendanceRule, error) {
	var rules []model.StudentAttendanceRule
	err := config.DB.Find(&rules).Error
	return rules, err
}

func GetStudentAttendanceRuleByID(id uint) (model.StudentAttendanceRule, error) {
	var rule model.StudentAttendanceRule
	err := config.DB.First(&rule, id).Error
	return rule, err
}

func GetStudentAttendanceRuleByDay(day string) (model.StudentAttendanceRule, error) {
	var rule model.StudentAttendanceRule
	err := config.DB.Where("day = ?", day).Find(&rule).Error
	return rule, err
}
