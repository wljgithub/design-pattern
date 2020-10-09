package main

import "log"

func main() {
	walletFacede := newWalletfacade("abc",1234)
	err := walletFacede.addMoneyToWallet("abc",1234,10)
	if err != nil {
		log.Fatal(err)
	}

	err = walletFacede.deductMoneyFromWallet("ab",1234,5)
	if err != nil {
		log.Fatal(err)
	}

}
