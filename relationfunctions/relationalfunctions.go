package relationfunctions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Amizhthanmd/Golang_Postgresql_Relationships/models"
	"github.com/Amizhthanmd/Golang_Postgresql_Relationships/queryfunctions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

func SetupDB() {
	dsn := "host=localhost user=postgres password=Ami160320! dbname=sqlrelationships port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database : ", err)
	}
	fmt.Println("Databse connected successfully")
}

func MigrateDB() {
	err := db.AutoMigrate(&models.Employees{}, &models.Empaccountdetails{}, &models.Country{}, &models.City{}, &models.Username{}, &models.Skills{})
	if err != nil {
		log.Fatal("failed to migrate : ", err)
	}
	fmt.Println("DB migrated successfully")
}

// One-to-One Relationships :-

func CreateEmployee(c *gin.Context) {
	var employee models.Employees

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	employeeDB := queryfunctions.NewPostgresqlDB(db)

	if err := employeeDB.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employees"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employeedetails created successfully", "user": employee})
}

func GetEmployeeWithDetails(c *gin.Context) {

	employeeDB := queryfunctions.NewPostgresqlDB(db)

	employees, err := employeeDB.GetEmployeeWithDetails()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employees": employees})
}

func GetAccEmployees(c *gin.Context) {

	accempDB := queryfunctions.NewPostgresqlDB(db)

	accemployees, err := accempDB.GetAccEmployees()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "AccountsWithEmployee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"AccountsWithEmployee" : accemployees})
}

// One-to-Many Relationships :-

func CreateCountryCity(c *gin.Context) {
	var country models.Country

	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	countryDB := queryfunctions.NewPostgresqlDB(db)

	if err := countryDB.CreateCountryCity(&country); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create countrycity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CountryCity created successfully", "user": country})
}

func GetCountryCity(c *gin.Context) {

	countryDB := queryfunctions.NewPostgresqlDB(db)

	countries, err := countryDB.GetCountryCity()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"CountrywithCity": countries})
}

func GetCityCountry(c *gin.Context) {

	cityDB := queryfunctions.NewPostgresqlDB(db)

	cities, err := cityDB.GetCityCountry()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"CitywithCountry": cities})

}

// Many-to-Many Relationships :-

func CreateUsernameSkills(c *gin.Context) {
	var username models.Username

	if err := c.ShouldBindJSON(&username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	usernameDB := queryfunctions.NewPostgresqlDB(db)

	if err := usernameDB.CreateUsernameSkills(&username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create username&skills"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UsernameSkills created successfully", "user": username})
}

func GetUersSkills(c *gin.Context) {

	usernameskillsDB := queryfunctions.NewPostgresqlDB(db)

	users, err := usernameskillsDB.GetUersSkills()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UsernameSkills not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"UsernameWithSkills": users})

}

func GetSkillsUser(c *gin.Context) {

		skillsuserDB := queryfunctions.NewPostgresqlDB(db)

		skills, err := skillsuserDB.GetSkillsUser()

		if err != nil{
			c.JSON(http.StatusNotFound, gin.H{"error":"Skillsuser not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"SkillsWithUsernames" : skills})
}