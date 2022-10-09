package main

import (
	"fmt"
	"sync"
	"time"
)

type WriteData struct {
	Data string
	Time time.Duration
}

func DifficultFuncChanStart() {
	var chWrite = make(chan WriteData, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			defer close(chWrite)
			wg.Done()
		}()
		DifficultFuncChanWrite(chWrite)
	}()

	go func() {
		defer wg.Done()
		DifficultFuncChanRead(chWrite)
	}()

	wg.Wait()
}

func DifficultFuncChanWrite(chWrite chan<- WriteData) {
	var data = []WriteData{
		{Data: "first", Time: 1 * time.Second},
		{Data: "second", Time: 2 * time.Second},
		{Data: "third", Time: 3 * time.Second},
		{Data: "fourth", Time: 4 * time.Second},
		{Data: "fifth", Time: 5 * time.Second},
	}

	for _, data := range data {
		select {
		case chWrite <- data:
		case <-time.After(3 * time.Second):
			fmt.Println("chWrite timeout expire")
			return
		}
	}
}

func DifficultFuncChanRead(chWrite <-chan WriteData) {
	for data := range chWrite {
		time.Sleep(data.Time)
		fmt.Printf("DifficultFuncChan: %s\r\n", data.Data)
	}
}
