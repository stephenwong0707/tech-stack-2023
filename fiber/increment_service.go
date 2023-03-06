package main

import (
	"fmt"
	"time"
)

type IncrementSevice struct {
	list []int64
	ch   chan int64
}

func (s *IncrementSevice) in() {
	for {
		select {
		case c := <-s.ch:
			s.list = append(s.list, c)
		}
	}
}

func main() {
	incrementSevice := &IncrementSevice{
		list: []int64{},
		ch:   make(chan int64),
	}

	go func() {
		for i := 0; i < 500; i++ {
			fmt.Println("gorountine")
			incrementSevice.ch <- int64(i)
		}
	}()
	go incrementSevice.in()

	for i := 0; i < 499; i++ {
		fmt.Println("main")
		incrementSevice.ch <- int64(i)
	}
	time.Sleep(2 * time.Second)

	fmt.Printf("main %d ", len(incrementSevice.list))
}
