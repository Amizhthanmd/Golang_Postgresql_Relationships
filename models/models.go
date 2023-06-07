package models

import (
	"time"

	"gorm.io/gorm"
)

// One-to-One Relationship

type Employees struct {
	Empid             uint `gorm:"unique"`
	Name              string
	Age               int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt    `gorm:"index"`
	Empaccountdetails *Empaccountdetails `gorm:"foreignKey:Empid"`
}

type Empaccountdetails struct {
	Empid      uint `gorm:"primaryKey;unique"`
	Accountnum int  `gorm:"unique"`
	Phonenum   int  `gorm:"unique"`
	Email      *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Employees  *Employees     `gorm:"foreignKey:Empid"`
}

// One-to-Many Relationship

type Country struct {
	Cid         uint   `gorm:"primaryKey;unique"`
	Countryname string `gorm:"unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Cities      []City         `gorm:"foreignKey:CountryID"`
}

type City struct {
	Cityid     uint   `gorm:"primaryKey"`
	Cityname   string `gorm:"unique"`
	Population int
	CountryID  uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Country    *Country       `gorm:"foreignKey:CountryID"`
}

// Many-to-Many Relationship

type Username struct {
	Userid     uint `gorm:"primaryKey"`
	Username   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Userskills []Skills `gorm:"many2many:user_skills;"`
}

type Skills struct {
	Skid        uint   `gorm:"primaryKey"`
	Skill       string `gorm:"unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Skillsusers []Username `gorm:"many2many:user_skills;"`
}
