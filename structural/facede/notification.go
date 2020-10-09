package main

import "fmt"

type notification struct {

}

func newNotification() *notification {
	return &notification{}
}
func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification)sendWalletDebitNotification()  {
	fmt.Println("Sending wallet debit notification")
}
