package main

import "fmt"

type customer struct {
	id string
}

func (this *customer)update(itemName string)  {
	fmt.Printf("Sending email to customer %s for item %s\n",this.id,itemName)
}

func (this *customer) getID() string {
	return this.id
}
