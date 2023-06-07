package queryfunctions

import (
	"github.com/Amizhthanmd/Golang_Postgresql_Relationships/models"
	"gorm.io/gorm"
)

type PostgresqlDB struct {
	db *gorm.DB
}

func NewPostgresqlDB(db *gorm.DB) *PostgresqlDB {
	return &PostgresqlDB{db: db}
}

// One-to-One Relationships :-

func (e *PostgresqlDB) CreateEmployee(employee *models.Employees) error {
	if err := e.db.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

func (e *PostgresqlDB) GetEmployeeWithDetails() ([]models.Employees, error) {
	var employees []models.Employees
	if err := e.db.Model(&models.Employees{}).Preload("Empaccountdetails").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (e *PostgresqlDB) GetAccEmployees()  ([]models.Empaccountdetails, error){
	var accemployees []models.Empaccountdetails
	if err := e.db.Model(&models.Empaccountdetails{}).Preload("Employees").Find(&accemployees).Error; err != nil {
		return nil, err
	}
	return accemployees, nil
}

// One-to-Many Relationships :-

func (e *PostgresqlDB) CreateCountryCity(countrycity *models.Country) error {
	if err := e.db.Create(countrycity).Error; err != nil {
		return err
	}
	return nil
}

func (e *PostgresqlDB) GetCountryCity() ([]models.Country, error) {
	var country []models.Country
	if err := e.db.Model(&models.Country{}).Preload("Cities").Find(&country).Error; err != nil {
		return nil, err
	}
	return country, nil
}

func (e *PostgresqlDB) GetCityCountry() ([]models.City, error) {
	var city []models.City
	if err := e.db.Model(&models.City{}).Preload("Country").Find(&city).Error; err != nil {
		return nil, err
	}
	return city, nil
}

// Many-to-Many Relationships :-

func (e *PostgresqlDB) CreateUsernameSkills(username *models.Username) error {
	if err := e.db.Create(username).Error; err != nil {
		return err
	}
	return nil
}

func (e *PostgresqlDB) GetUersSkills() ([]models.Username, error) {
	var userskills []models.Username
	if err := e.db.Model(&models.Username{}).Preload("Userskills").Find(&userskills).Error; err != nil {
		return nil, err
	}
	return userskills, nil
}

func (e *PostgresqlDB) GetSkillsUser() ([]models.Skills, error){
	var Skillsuser []models.Skills
	if err:= e.db.Model(&models.Skills{}).Preload("Skillsusers").Find(&Skillsuser).Error; err != nil {
		return nil, err
	}
	return Skillsuser, nil
}