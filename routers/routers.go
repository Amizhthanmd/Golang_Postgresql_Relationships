package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Amizhthanmd/Golang_Postgresql_Relationships/relationfunctions"
)

func Routers() {
	relationfunctions.SetupDB()
	relationfunctions.MigrateDB()

	r := gin.Default()

	r.POST("/Createemployees", relationfunctions.CreateEmployee)
	r.GET("/Getempdetails", relationfunctions.GetEmployeeWithDetails)
	r.GET("/Getaccemployees", relationfunctions.GetAccEmployees)

	r.POST("/Createcountrycity", relationfunctions.CreateCountryCity)
	r.GET("/Getcountrycity", relationfunctions.GetCountryCity)
	r.GET("/Getcitycountry",relationfunctions.GetCityCountry)

	r.POST("/Createusernameskills", relationfunctions.CreateUsernameSkills)
	r.GET("/Getuersskills", relationfunctions.GetUersSkills)
	r.GET("/Getskillsuser", relationfunctions.GetSkillsUser)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start...")
	}
}


