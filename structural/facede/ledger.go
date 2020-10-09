package main

import "fmt"

type ledger struct {

}

func newLedger() *ledger {
	return &ledger{}
}
func (s *ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId [%s] with txnType %s for amount %d\n",accountId,txnType,amount)
}