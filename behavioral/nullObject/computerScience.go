package main

type computerScience struct {
	numberOfProfessors int
}


func (this *computerScience)getNumerOfProfessors()int  {
	return this.numberOfProfessors
}

func (this *computerScience) getName() string {
	return "computerscience"
}
