package main

import (
	"net/http"
	"strconv"
	"log"

	"github.com/labstack/echo/v4"
)


type IncrementCounterReq struct {
	ToIncrementBy int `json:"toIncrementBy"`
}



func AddIncrements(c echo.Context) error {
	toIncrementByReq := new(IncrementCounterReq)
	if err := c.Bind(toIncrementByReq); err != nil {
		return err
	}

	log.Printf("Incremented by: %d\n", toIncrementByReq.ToIncrementBy)

	counter.IncrementCounter(toIncrementByReq.ToIncrementBy)
	return c.String(http.StatusCreated, strconv.Itoa(counter.value + toIncrementByReq.ToIncrementBy))
}

func GetIncrements(c echo.Context) error {
	return c.String(http.StatusOK, strconv.Itoa(counter.value))
}



func ResetCounter(c echo.Context) error {
	return c.String(http.StatusOK, strconv.Itoa(counter.ResetValue()))
}