package model

import (
	"time"
)

type Student struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Photo           *string    `json:"photo"` // bisa null
	Name            string     `json:"name" validate:"required"`
	NISN            string     `json:"nisn" validate:"required"`
	NIK             *string    `json:"nik"`
	BirthPlace      *string    `json:"birth_place"`
	BirthDate       *time.Time `json:"birth_date"` // nullable date
	Gender          *string    `json:"gender" gorm:"type:enum('laki-laki','perempuan')"`
	Address         *string    `json:"address"`
	PhoneNumber     *string    `json:"no_telp"`
	Email           *string    `json:"email" validate:"omitempty,email"`
	KebutuhanKhusus bool       `json:"kebutuhan_khusus"` // tinyint(1) -> bool
	Disabilitas     bool       `json:"disabilitas"`
	FatherName      *string    `json:"father_name"`
	MotherName      *string    `json:"mother_name"`
	UserID          *uint64    `json:"user_id"` // foreign key
	Status          bool       `json:"status"`  // tinyint(1), default 1
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
