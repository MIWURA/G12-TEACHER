package entity

import (
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("project-teacher.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Officer{},
		&Faculty{},
		&Prefix{},
		&Educational{},
		&Teacher{},
	)

	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)

	Officer1 := Officer{
		Name:     "panupol",
		Email:    "panupol@gamil.com",
		Password: string(password),
	}
	db.Model(&Officer{}).Create(&Officer1)

	Officer2 := Officer{
		Name:     "somtum",
		Email:    "somtum@gamil.com",
		Password: string(password),
	}
	db.Model(&Officer{}).Create(&Officer2)

	Science := Faculty{
		Name: "สำนักวิชาวิทยาศาสตร์",
	}
	db.Model(&Faculty{}).Create(&Science)

	agricultural_technology := Faculty{
		Name: "สำนักวิชาเทคโนโลยีการเกษตร",
	}
	db.Model(&Faculty{}).Create(&agricultural_technology)

	social_technology := Faculty{
		Name: "สำนักวิชาเทคโนโลยีสังคม",
	}
	db.Model(&Faculty{}).Create(&social_technology)

	engineering := Faculty{
		Name: "สำนักวิชาวิศวกรรมศาสตร์",
	}
	db.Model(&Faculty{}).Create(&engineering)

	medicine := Faculty{
		Name: "สำนักวิชาแพทย์ศาสตร์",
	}
	db.Model(&Faculty{}).Create(&medicine)

	Nursing := Faculty{
		Name: "สำนักวิชาพยาบาลศาสตร์",
	}
	db.Model(&Faculty{}).Create(&Nursing)

	Prefix_1 := Prefix{
		Name: "อ.",
	}
	db.Model(&Prefix{}).Create(&Prefix_1)

	Prefix_2 := Prefix{
		Name: "อ. ดร.",
	}
	db.Model(&Prefix{}).Create(&Prefix_2)

	Prefix_3 := Prefix{
		Name: "ศ. ดร.",
	}
	db.Model(&Prefix{}).Create(&Prefix_3)

	Prefix_4 := Prefix{
		Name: "ผศ. ดร.",
	}
	db.Model(&Prefix{}).Create(&Prefix_4)

	Prefix_5 := Prefix{
		Name: "รศ. ดร.",
	}
	db.Model(&Prefix{}).Create(&Prefix_5)

	Prefix_6 := Prefix{
		Name: "Asst. Prof.",
	}
	db.Model(&Prefix{}).Create(&Prefix_6)

	Prefix_7 := Prefix{
		Name: "Assoc. Prof.",
	}
	db.Model(&Prefix{}).Create(&Prefix_7)

	Prefix_8 := Prefix{
		Name: " Prof.",
	}
	db.Model(&Prefix{}).Create(&Prefix_8)

	Educational_1 := Educational{
		Name: "ปริญญาโท",
	}
	db.Model(&Educational{}).Create(&Educational_1)

	Educational_2 := Educational{
		Name: "ปริญญาเอก",
	}
	db.Model(&Educational{}).Create(&Educational_2)

	db.Model(&Teacher{}).Create(&Teacher{
		Name:        "data",
		Email:       "data@gmail.com",
		Password:    "dt_12@4321",
		Officer:     Officer1,
		Faculty:     engineering,
		Prefix:      Prefix_3,
		Educational: Educational_2,
	})

	db.Model(&Teacher{}).Create(&Teacher{
		Name:        "botton",
		Email:       "botton@gmail.com",
		Password:    "bottom1234_12",
		Officer:     Officer1,
		Faculty:     engineering,
		Prefix:      Prefix_3,
		Educational: Educational_2,
	})

	db.Model(&Teacher{}).Create(&Teacher{
		Name:        "tiger",
		Email:       "tiger@gmail.com",
		Password:    "t123451231",
		Officer:     Officer2,
		Faculty:     social_technology,
		Prefix:      Prefix_3,
		Educational: Educational_2,
	})

	var panupol Officer
	var somtum Officer
	db.Raw("SELETE * FROM officer WHERE email = ?", "panupol@gamil.com").Scan(&panupol)
	db.Raw("SELETE * FROM officer WHERE email = ?", "somtum@gamil.com").Scan(&somtum)

}
