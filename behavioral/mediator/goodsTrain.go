package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (this *goodsTrain)requestArrival()  {
	if this.mediator.canLand(this){
		fmt.Println("GoodsTrain: landing")
	}else{
		fmt.Println("GoodsTrain: waiting")
	}
}

func (this *goodsTrain) departure() {
	this.mediator.notifyFree()
	fmt.Println("GoodsTrain: leaving")
}

func (this *goodsTrain) permitArrival() {
	fmt.Println("GoodsTrain: arrival permitted. landing")
}
