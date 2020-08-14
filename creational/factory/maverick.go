package main

type mavarick struct {
	gun
}

func newMaverick() iGun {
	mavarick := &gun{}
	mavarick.setName("mavarick")
	mavarick.setPower(99)
	return mavarick
}
