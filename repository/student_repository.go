package repository

import (
	"github.com/farinchan/thesis-attendance-backend/config"
	"github.com/farinchan/thesis-attendance-backend/model"
)

func GetAllStudent() ([]model.Student, error) {
	var students []model.Student
	err := config.DB.Find(&students).Error
	return students, err
}

func GetStudentByNisn(nisn string) (model.Student, error) {
	var student model.Student
	err := config.DB.Where("nisn = ?", nisn).First(&student).Error
	return student, err
}
