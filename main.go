package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/stundents", getStudents)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
 
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "list of all students")
}