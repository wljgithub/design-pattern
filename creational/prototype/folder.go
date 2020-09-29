package main

import "fmt"

type folder struct {
	childrents []inode
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name + "_clone")
	for _,i := range f.childrents{
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name}
	var tempChildrens []inode

	for _,i:=range f.childrents{
		copy := i.clone()
		tempChildrens = append(tempChildrens,copy)
	}
	cloneFolder.childrents = tempChildrens
	return cloneFolder
}
