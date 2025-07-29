package repository

import (
	"github.com/farinchan/thesis-attendance-backend/config"
	"github.com/farinchan/thesis-attendance-backend/model"
)

func GetStudentAttendanceByStudentIdAndDate(studentId uint64, date string) (model.StudentAttendance, error) {
	var attendance model.StudentAttendance
	err := config.DB.Where("student_id = ? AND date = ?", studentId, date).First(&attendance).Error
	return attendance, err
}

func CreateStudentAttendance(attendance *model.StudentAttendance) error {
	return config.DB.Create(attendance).Error
}
