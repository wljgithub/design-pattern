package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (this *passengerTrain) requestArrival() {
	if this.mediator.canLand(this){
		fmt.Println("PassengerTrain: Landing")
	}else{
		fmt.Println("PassengerTrain: Waiting")
	}

}

func (this *passengerTrain) departure() {
	fmt.Println("PassengerTrain: Leaving")
	this.mediator.notifyFree()
}

func (this *passengerTrain)permitArrival()  {
	fmt.Println("PassengerTrain: Arrival permitted. landing")
}

