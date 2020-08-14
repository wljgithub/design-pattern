package main

type ak47 struct {
	gun
}

func newAK47() iGun {
	ak47 := &gun{}
	ak47.setName("ak47")
	ak47.setPower(100)

	return ak47
}
