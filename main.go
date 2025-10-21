package main

import (
	"net/http"
   "github.com/leonardoHPalves/api-students/db"


  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"

)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/stundents", getStudents)
  e.POST("/students", createStudent) 
  e.GET("/students/:id", getStudent)
  e.PUT("/students/:id", updateStudent)
  e.DELETE("/students/:id", deleteStudent)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
 
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "list of all students")
}

func createStudent(c echo.Context) error {
  student := db.Student{}
   if err := c.Bind(&student); err != nil{
      return err
   }
	if err := db.AddStudent(student); err != nil {
    return c.String(http.StatusInternalServerError, "Error to create student")
  }
  return c.String(http.StatusOK, "create student")
}

func getStudent(c echo.Context) error {
  return c.String(http.StatusOK, "get student")
}

func updateStudent(c echo.Context) error {
  return c.String(http.StatusOK, "update student")
}

func deleteStudent(c echo.Context) error {
  return c.String(http.StatusOK, "delete students")
}