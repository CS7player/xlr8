package main

import "os"

func GetCurrentDir() (dir string, err error) {
	dir, err = os.Getwd()
	if err != nil {
		return dir, err
	}
	return dir, err
}

func GetFoldersList(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []os.DirEntry{}
	}
	return entries
}
