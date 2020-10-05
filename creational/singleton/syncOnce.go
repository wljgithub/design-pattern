package main

import (
	"fmt"
	"sync"
)

var once = new(sync.Once)

type single struct {

}

var singleInstance *single

func getInstance()*single{
	if singleInstance == nil{
		once.Do(func() {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		})
	}else {
		fmt.Println("Single instance already created-2")
	}
	return singleInstance
}

func main()  {
	for i:=0;i<100;i++{
		go getInstance()
	}
	fmt.Scanln()
}
