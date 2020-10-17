package main

type college struct {
	departments []department
}

func (this *college)addDepartment(departmentName string,numOfProfessors int)  {
	if departmentName == "computerscience"{
		computerScienceDepartment := &computerScience{numberOfProfessors: numOfProfessors}
		this.departments = append(this.departments,computerScienceDepartment)
	}
	if departmentName == "mechanical"{
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		this.departments = append(this.departments,mechanicalDepartment)
	}
	return
}

func (this *college)getDepartment(departmentName string)department  {
	for _,department := range this.departments{
		if department.getName() == departmentName{
			return department
		}
	}
	return &nullDepartment{}
}
