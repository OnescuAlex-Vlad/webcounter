package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"
	"log"

	"github.com/labstack/echo/v4"
)

var wg sync.WaitGroup
var counter = NewCounter()

type Counter struct {
	value int
	ch    chan func()
}

func NewCounter() Counter {
	counter := &Counter{0, make(chan func(), 100)}
	go func(counter *Counter) {
		for f := range counter.ch {
			f()
		}
	}(counter)
	return *counter
}

type IncrementCounterReq struct {
	ToIncrementBy int `json:"toIncrementBy"`
}

func (counter *Counter) IncrementEverySecond() {
	
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		counter.value++
		log.Printf("The value is: %d\n", counter.value)
	}

}

func (counter *Counter) IncrementCounter(toIncrementBy int) {
	counter.ch <- func() {
		for i := 0; i < toIncrementBy; i++ {
			wg.Add(1)
			counter.value++
			wg.Done()
		}
	}

	wg.Wait()
}

func (counter *Counter) GetCounterValue() int {
	ret := make(chan int)

	wg.Add(1)
	counter.ch <- func() {
		ret <- counter.value
		wg.Done()
		close(ret)
	}

	wg.Wait()
	return <-ret
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

func (counter *Counter) ResetValue() int {
	ret := make(chan int)

	counter.value = 0

	wg.Add(1)
	counter.ch <- func() {
		
		ret <- counter.value
		wg.Done()
		close(ret)
	}

	wg.Wait()

	return counter.value
}

func ResetCounter(c echo.Context) error {
	return c.String(http.StatusOK, strconv.Itoa(counter.ResetValue()))
}


func main() {

	go counter.IncrementEverySecond()

	e := echo.New()
	e.GET("/counter", GetIncrements)
	e.POST("/counter", AddIncrements)
	e.POST("/counter/reset", ResetCounter)


	e.Logger.Fatal(e.Start(":1323"))
}
