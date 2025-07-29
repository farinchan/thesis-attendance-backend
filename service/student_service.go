package service

import (
	"github.com/farinchan/thesis-attendance-backend/model"
	"github.com/farinchan/thesis-attendance-backend/repository"
)

func GetStudents() ([]model.Student, error) {
	return repository.GetAllStudent()
}

func GetStudentByNisn(nisn string) (model.Student, error) {
	return repository.GetStudentByNisn(nisn)
}
