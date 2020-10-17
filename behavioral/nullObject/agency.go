package main

import (
	"fmt"
)

func main() {
	college1 := createCollege1()
	college2 := createCollege2()

	totalProfessor := 0
	departmentArray := []string{"computerscience","mechanical","civil","electronics"}

	for _,departmentName := range departmentArray{
		d := college1.getDepartment(departmentName)
		totalProfessor += d.getNumerOfProfessors()
	}

	fmt.Printf("Total number of professor in college1 is %d\n",totalProfessor)

	totalProfessor =0
	for _,departmentName := range departmentArray{
		d := college2.getDepartment(departmentName)
		totalProfessor += d.getNumerOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n",totalProfessor)
}

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience",4)
	college.addDepartment("mechanical",5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerscience",2)
	return college
}
