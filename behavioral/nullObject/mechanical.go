package main

type mechanical struct {
	numberOfProfessors int
}

func (this *mechanical) getNumerOfProfessors() int {
	return this.numberOfProfessors
}


func (this *mechanical) getName() string {
	return "mechanical"
}