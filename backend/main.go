package main

import (
	"github.com/MIWURA/Project-SA-teacher/entity"
	"github.com/MIWURA/Project-SA-teacher/controller"
	"github.com/MIWURA/Project-SA-teacher/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("/")

{
	router.Use(middlewares.Authorizes())
	{
		// Teacher
		r.GET("/educationnals", controller.ListEducational)
		r.GET("/educationnals/:id", controller.GetEducational)
		r.POST("/educationnals", controller.CreateEducational)
		r.PATCH("/educationnals", controller.UpdateEducational)
		r.DELETE("/educationnals/:id", controller.DeleteEducational)

		//Student
		r.GET("/faculties", controller.ListFaculty)
		r.GET("/faculties/:id", controller.GetFaculty)
		r.POST("/faculties", controller.CreateFaculty)
		r.PATCH("/faculties", controller.UpdateFaculty)
		r.DELETE("/faculties/:id", controller.DeleteFaculty)

		//Teacher
		r.GET("/teachers", controller.ListTeachers)
		r.GET("/teachers/:id", controller.GetTeacher)
		r.POST("/teachers", controller.CreateTeacher)
		r.PATCH("/teachers", controller.UpdateTeacher)
		r.DELETE("/teachers/:id", controller.DeleteTeacher)

		//Teaching_duration
		r.GET("/officers", controller.ListOfficers)
		r.GET("/officers/:id", controller.GetOfficer)
		r.POST("/officers", controller.CreateOfficer)
		r.PATCH("/officers", controller.UpdateOfficer)
		r.DELETE("/officers/:id", controller.DeleteOfficer)

		//Content_difficulty_level
		r.GET("/prefixes", controller.ListPrefix)
		r.GET("/prefixes/:id", controller.GetPrefix)
		r.POST("/prefixes", controller.CreatePrefix)
		r.PATCH("/prefixes", controller.UpdatePrefix)
		r.DELETE("/prefixes/:id", controller.DeletePrefix)
	// Run the server
		}

}

	r.POST("/signup",  controller.CreateOfficer)
	// login User Route
	r.POST("/login", controller.Login)

	// // Run the server go run main.go
	r.Run()



}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}

}
