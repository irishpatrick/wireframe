package main

import (
	"os"
	"path/filepath"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fn string) (string, error) {
	dir, err := filepath.Abs(fn)
	f, err := os.Open(dir)
	check(err)

	stat, err := f.Stat()
	check(err)

	buffer := make([]byte, stat.Size()+1)
	_, err = f.Read(buffer)
	check(err)

    buffer[stat.Size()] = 0x00

	f.Close()

	return string(buffer), nil
}
