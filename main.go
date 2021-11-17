package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var eachSecond int
	flag.IntVar(&eachSecond, "each", 5, "wait each sec")
	var untilSecond int
	flag.IntVar(&untilSecond, "until", 10, "until sec")
	flag.Parse()

	var wg sync.WaitGroup
	stop := func() <-chan interface{} {
		wg.Add(1)
		s := make(chan interface{})
		go func() {
			defer func() {
				wg.Done()
			}()
			defer func() {
				close(s)
			}()
			time.Sleep(time.Duration(untilSecond) * time.Second)
		}()
		return s
	}()

	wg.Add(1)
	go func() {
		defer func() { wg.Done() }()
		ticker := time.NewTicker(time.Duration(eachSecond)*time.Second + 100*time.Millisecond)
		defer func() { ticker.Stop() }()
	loop:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			case <-stop:
				break loop
			}
		}
	}()

	wg.Wait()
	fmt.Println("done")
}
