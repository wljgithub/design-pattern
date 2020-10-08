package main

import "fmt"

// component.go
type component interface {
	search(string)
}

// folder.go

type folder struct {
	components []component
	name string
}

func (f * folder)search(keyword string)  {
	fmt.Printf("Search recursively for keyword %s in folder %s\n",keyword,f.name)
	for _, composite := range f.components{
		composite.search(keyword)
	}
}

func (f *folder)add(c component)  {
	f.components = append(f.components,c)
}

// file.go

type file struct {
	name string
}

func (f *file)search(keyword string)  {
	fmt.Printf("Searching for keyword %s in file %s\n",keyword,f.name)
}

func (f *file) getName() string {
	return f.name
}

func main() {
	file1:= &file{name: "File1"}
	file2:= &file{name: "File1"}
	file3:= &file{name: "File3"}
	folder1 := &folder{name: "Folder1"}
	folder1.add(file1)

	folder2 := folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	folder2.search("rose")
}
