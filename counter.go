package main

import (
	"time"
	"sync"
	"log"
)

var wg sync.WaitGroup

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