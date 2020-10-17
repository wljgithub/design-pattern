package main

type nullDepartment struct {
	numberOfProfessors int
}

func (this *nullDepartment) getNumerOfProfessors() int {
	return this.numberOfProfessors
}


func (this *nullDepartment)getName() string  {
	return "nullDepartment"
}
