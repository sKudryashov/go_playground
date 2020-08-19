package main

import (
	"fmt"
	"strings"
)

func main() {
	fs := NewFileSystem()
	fs.MkDir("/foo") //-> succed
	// fs.MkDir("/foo/bar/foo") //-> fail;
	// fs.MkDir("/foo/bar") //-> fail;
	// fs.MkDir("/foo/bar") // - succeed
	// fs.MkDir("/foo/bar/foo") // succeed

	// fs.MkDir("/foo/bar/foo/bar")
	// fs.WriteFile("/foo/bar/foo/bar/file.txt")
	// fs.ReadFile("/foo/bar/foo/bar/file.txt")
	// fs.ReadFile("/foo/bar/foo/file.txt") // error
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		dirPaths: make(map[string]map[string]string),
		// filepath: make(map[string]string),
	}
}

type FileSystem struct {
	dirPaths map[string]map[string]string
}

func (fs *FileSystem) MkDir(dir string) error {
	if _, ok := fs.dirPaths[dir]; ok {
		return fmt.Errorf("the dir %s already exists", dir)
	}
	dir = strings.Trim(dir, "/")
	println(" dir ", dir)
	dirPathArr := strings.Split(dir, "/")
	basePathArr := dirPathArr[:len(dirPathArr)-1]
	basePath := strings.Join(basePathArr, "/")

	println(" basePath ", basePath)
	fmt.Println(" fs.dirPaths ", fs.dirPaths)
	if _, ok := fs.dirPaths[basePath]; !ok {
		fs.dirPaths[dir] = nil
		println(" the dir created ", dir)
	} else {
		println("wrong ierarchy")
		return fmt.Errorf("wrong ierarchy")
	}
	return nil
}

func (fs *FileSystem) ReadFile(path string) (string, error) {
	pathsSplit := strings.Split(path, "/")
	fileName := pathsSplit[len(pathsSplit)-1]
	dirPathArr := pathsSplit[0 : len(pathsSplit)-2]
	dirPath := strings.Join(dirPathArr, "/")
	println(" dirPath ", path)
	if file, ok := fs.dirPaths[dirPath]; ok {
		if content, ok := file[fileName]; !ok {
			return "", fmt.Errorf("file is not found")
		} else {
			return content, nil
		}
	} else {
		return "", fmt.Errorf("unable to find dir")
	}
}

func (fs *FileSystem) WriteFile(path string, content string) error {
	//
	return nil
}
