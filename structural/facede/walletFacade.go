package main

import "fmt"

type walletFacade struct {
	account *account
	wallet *wallet
	securityCode *securityCode
	notification *notification
	ledger *ledger
}

func newWalletfacade(accountId string,code int) *walletFacade  {
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account: newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet:newWallet(),
		notification: newNotification(),
		ledger: newLedger(),
	}
	fmt.Println("account created")
	return walletFacade
}

func (w *walletFacade)addMoneyToWallet(accountId string,securityCode int,amount int)error  {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountId)
	if err !=nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId,"credit",amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountId string, securityCode, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountId)

	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId,"debit",amount)
	return nil
}
