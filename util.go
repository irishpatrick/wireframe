package main

import (
	"errors"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func itob(i int) bool {
	if i != 0 {
		return true
	}
	return false
}

func parseSlice(s []float32, index int, points int) ([]float32, error) {
	out := make([]float32, 0)

	ptr := index * points
	end := ptr + points
	if end > len(s) {
		return nil, errors.New("parseSlice: index out of range")
	}

	for i := ptr; i < end; i++ {
		out = append(out, s[i])
	}

	return out, nil
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
