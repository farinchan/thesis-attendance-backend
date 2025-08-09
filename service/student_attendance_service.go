package service

import (
	"github.com/farinchan/thesis-attendance-backend/model"
	"github.com/farinchan/thesis-attendance-backend/repository"
)

func GetStudentAttendanceBystudentIdAndDate(studentId uint64, date string) (model.StudentAttendance, error) {
	return repository.GetStudentAttendanceByStudentIdAndDate(studentId, date)
}

func CreateStudentAttendance(attendance *model.StudentAttendance) error {
	return repository.CreateStudentAttendance(attendance)
}

func UpdateStudentAttendance(attendance *model.StudentAttendance) error {
	return repository.UpdateStudentAttendance(attendance)
}
