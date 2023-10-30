package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-echo-gorm-practice/controller"
	"go-echo-gorm-practice/storage"
)

func main() {
  var env *string = flag.String("env", "local", "environment variable")
  var schema *string = flag.String("schema", "test", "db schema name")

  flag.Parse()

  if *env == "local" {
    // local
    fmt.Println("local")
  } else if *env == "dev" {
    // dev
    fmt.Println("dev")
  } else if *env == "prod" {
    // prod
    fmt.Println("prod")
  }
  
  // Echo instance
  e := echo.New()

  storage.NewDB(*schema)

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  e.GET("/students", controller.GetStudents)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}