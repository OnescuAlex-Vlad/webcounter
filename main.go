package main

import (
	"github.com/labstack/echo/v4"
)


var counter = NewCounter()


func main() {

	go counter.IncrementEverySecond()

	e := echo.New()
	e.GET("/counter", GetIncrements)
	e.POST("/counter", AddIncrements)
	e.POST("/counter/reset", ResetCounter)


	e.Logger.Fatal(e.Start(":1323"))
}
