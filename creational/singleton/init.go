package main

import "fmt"

func init()  {
	singleInstance = &single{}
}

type single struct {

}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		fmt.Println("Instance is not created yet")
	}else {
		fmt.Println("Instance was created")
	}
	return singleInstance

}
func main()  {
	for i:=0;i<100;i++{
		go getInstance()
	}
	fmt.Scanln()
}
