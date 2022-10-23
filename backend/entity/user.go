package entity

import (
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Teacher  []Teacher `gorm:"ForeignKey:OfficerID"`
}

type Faculty struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Teacher []Teacher `gorm:"ForeignKey:FacultyID"`
}

type Prefix struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Teacher []Teacher `gorm:"ForeignKey:PrefixID"`
}

type Educational struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Teacher []Teacher `gorm:"ForeignKey:EducationalID"`
}

type Teacher struct {
	gorm.Model
	Name          string
	Email         string `gorm:"uniqueIndex"`
	Password      string
	OfficerID     *uint
	Officer       Officer `gorm:"references:id"`
	FacultyID     *uint
	Faculty       Faculty `gorm:"references:id"`
	PrefixID      *uint
	Prefix        Prefix `gorm:"references:id"`
	EducationalID *uint
	Educational   Educational `gorm:"references:id"`
}
