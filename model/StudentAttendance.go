package model

import (
    "time"
)

type StudentAttendance struct {
    ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
    StudentID   uint64     `json:"student_id" validate:"required"`
    Date        time.Time  `json:"date" validate:"required"` // format: YYYY-MM-DD
    TimeIn      *time.Time `json:"time_in"`                  // waktu masuk (optional)
    TimeInInfo  *string    `json:"time_in_info"`             // info tambahan masuk
    TimeOut     *time.Time `json:"time_out"`                 // waktu keluar (optional)
    TimeOutInfo *string    `json:"time_out_info"`            // info tambahan keluar
    TeacherID   *uint64    `json:"teacher_id"`               // bisa null
    CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
