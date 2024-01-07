package main

type File struct {
	name string
}

func (f *File) print(indentation string) {
	println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}
