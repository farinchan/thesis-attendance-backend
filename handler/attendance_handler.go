package handler

import (
	"github.com/farinchan/thesis-attendance-backend/model"
	"github.com/farinchan/thesis-attendance-backend/service"
	"github.com/farinchan/thesis-attendance-backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"time"
)

type nisn struct {
	NISN string `json:"nisn" validate:"required"`
}

var validate = validator.New()

func AttendanceCheckin(c *fiber.Ctx) error {

	var req nisn
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := validate.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	student, err := service.GetStudentByNisn(req.NISN)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = service.GetStudentAttendanceBystudentIdAndDate(student.ID, utils.GetCurrentDate())
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Attendance already recorded for today"})
	}

	hari := utils.GetIndonesianDayName(time.Now())
	rule, err := service.GetStudentAttendanceRulesByDay(hari)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Failed to retrieve attendance rule"})
	}

	var TimeInInfo string

	// Parse rule.Start (string) to time.Time
	startTime, err := time.Parse("15:04:00", rule.Start)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Invalid start time format"})
	}

	// Combine today's date with the parsed start time
	now := time.Now()
	startDateTime := time.Date(now.Year(), now.Month(), now.Day(), startTime.Hour(), startTime.Minute(), 0, 0, now.Location())

	if now.Before(startDateTime) {
		TimeInInfo = "Tepat Waktu"
	} else if now.After(startDateTime) {
		TimeInInfo = "Terlambat"
	} else {
		TimeInInfo = ""
	}
	// Ambil waktu saat ini
	currentTime := time.Now()

	timeInStr := currentTime.Format("15:04:05")
	attendance := model.StudentAttendance{
		StudentID:  student.ID,
		Date:       currentTime,
		TimeIn:     &timeInStr,
		TimeInInfo: &TimeInInfo,
		TeacherID:  func() *uint64 { id := uint64(1); return &id }(), // Assuming a default teacher ID, adjust as necessary
	}

	if err := service.CreateStudentAttendance(&attendance); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Failed to create attendance"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Attendance recorded successfully", "attendance": attendance})

}

func AttendanceCheckout(c *fiber.Ctx) error {

	var req nisn
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := validate.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	student, err := service.GetStudentByNisn(req.NISN)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	attendance, err := service.GetStudentAttendanceBystudentIdAndDate(student.ID, utils.GetCurrentDate())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Attendance not found for today"})
	}

	if attendance.TimeOut != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Attendance already checked out for today"})
	}

	hari := utils.GetIndonesianDayName(time.Now())
	rule, err := service.GetStudentAttendanceRulesByDay(hari)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Failed to retrieve attendance rule"})
	}

	var TimeOutInfo string

	// Parse rule.Start (string) to time.Time
	endTime, err := time.Parse("15:04:00", rule.End)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Invalid end time format"})
	}

	// Combine today's date with the parsed end time
	now := time.Now()
	endDateTime := time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), 0, 0, now.Location())
	if now.Before(endDateTime) {
		TimeOutInfo = "Pulang Cepat"
	} else if now.After(endDateTime) {
		TimeOutInfo = "Pulang Terlambat"
	} else {
		TimeOutInfo = ""
	}

	// Ambil waktu saat ini
	currentTime := time.Now()
	timeInStr := currentTime.Format("15:04:05")

	attendance.TimeOut = &timeInStr
	attendance.TimeOutInfo = &TimeOutInfo
	attendance.UpdatedAt = currentTime
	if err := service.UpdateStudentAttendance(&attendance); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + " - Failed to update attendance"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Attendance checked out successfully", "attendance": attendance})
}
