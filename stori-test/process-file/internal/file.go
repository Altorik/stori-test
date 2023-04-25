package internal

import "errors"

type File struct {
	Path string
	Name string
}

func (r *File) GetFullPath() string {
	if len(r.Path) < 1 {
		return r.Name
	}
	if r.Path[len(r.Path)-1:len(r.Path)] != "/" {
		return r.Path + "/" + r.Name
	}
	return r.Path + r.Name
}

func (r *File) NewFileToProcess(path, name string) error {
	if len(name) < 1 {
		return errors.New("invalid name")
	}
	if name[len(name)-4:len(name)] != ".csv" {
		return errors.New("invalid extension")
	}
	r.Name = name
	r.Path = path
	return nil
}
