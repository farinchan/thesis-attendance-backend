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
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := validate.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid NISN"})
	}


	student, err := service.GetStudentByNisn(req.NISN)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}

	_, err = service.GetStudentAttendanceBystudentIdAndDate(student.ID, utils.GetCurrentDate())
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Attendance already recorded for today"})
	}

	hari := utils.GetIndonesianDayName(time.Now())
	rule, err := service.GetStudentAttendanceRulesByDay(hari)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve attendance rules"})
	}

	var TimeInInfo string

	if time.Now().Before(rule.Start) {
		TimeInInfo = "Tepat Waktu"
	} else if time.Now().After(rule.Start) {
		TimeInInfo = "Terlambat"
	} else {
		TimeInInfo = ""
	}
	attendance := model.StudentAttendance{
		StudentID:  student.ID,
		Date:       time.Now(),
		TimeIn:     &time.Time{},
		TimeInInfo: &TimeInInfo,
	}

	if err := service.CreateStudentAttendance(&attendance); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to record attendance"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Attendance recorded successfully", "attendance": attendance})

}
