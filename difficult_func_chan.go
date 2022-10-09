package main

import (
	"fmt"
	"time"
)

type WriteData struct {
	Data string
	Time time.Duration
}

func DifficultFuncChanStart() {
	var chWrite = make(chan WriteData, 1)
	defer close(chWrite)

	go DifficultFuncChan(chWrite)
	for _, data := range []WriteData{
		{Data: "first", Time: 1 * time.Second},
		{Data: "second", Time: 2 * time.Second},
		{Data: "third", Time: 3 * time.Second},
		{Data: "fourth", Time: 4 * time.Second},
		{Data: "fifth", Time: 5 * time.Second},
	} {
		select {
		case chWrite <- data:
		case <-time.After(3 * time.Second):
			fmt.Println("chWrite timeout expire")
			return
		}
	}
}

func DifficultFuncChan(chWrite <-chan WriteData) {
	for data := range chWrite {
		time.Sleep(data.Time)
		fmt.Printf("DifficultFunc: %s\r\n", data.Data)
	}
}
