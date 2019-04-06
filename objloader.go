package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func printSlice(s []float32) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func iprintSlice(s []int32) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func loadMesh(fn string, program uint32) (Mesh, error) {
	mesh := Mesh{}

	fp, err := os.Open(fn)
	check(err)
	defer fp.Close()

	v := make([]float32, 0)
	vt := make([]float32, 0)
	n := make([]float32, 0)
	f := make([]int32, 0)
	buf := make([]float32, 0)

	// parse
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "v ") {
			parts := strings.Split(line, " ")
			for i := 1; i < 4; i++ {
				j, err := strconv.ParseFloat(parts[i], 32)
				check(err)
				v = append(v, float32(j))
			}
		}
		if strings.HasPrefix(line, "vt ") {
			parts := strings.Split(line, " ")
			for i := 1; i < len(parts); i++ {
				j, err := strconv.ParseFloat(parts[i], 32)
				check(err)
				vt = append(vt, float32(j))
			}
		}
		if strings.HasPrefix(line, "vn ") {
			parts := strings.Split(line, " ")
			for i := 1; i < len(parts); i++ {
				j, err := strconv.ParseFloat(parts[i], 32)
				check(err)
				n = append(n, float32(j))
			}
		}
		if strings.HasPrefix(line, "f ") {
			parts := strings.Split(line, " ")
			for i := 1; i < len(parts); i++ {
				references := strings.Split(parts[i], "/")
				for j := 0; j < len(references); j++ {
					k, err := strconv.ParseInt(references[j], 0, 32)
					check(err)
					f = append(f, int32(k))
				}
			}
		}
	}

	tick := 0
	for i := 0; i < len(f); i++ {
		if tick == 0 {
			index := f[i] - 1
			data, err := parseSlice(v, int(index), 3)
			check(err)
			buf = append(buf, data...)
		}
		if tick == 1 {
			index := f[i] - 1
			data, err := parseSlice(vt, int(index), 2)
			check(err)
			buf = append(buf, data...)
		}
		if tick == 2 {
			index := f[i] - 1
			data, err := parseSlice(n, int(index), 3)
			check(err)
			buf = append(buf, data...)
		}

		tick++
		if tick == 3 {
			tick = 0
		}
	}

	mesh.buffer = make([]float32, len(buf))
	mesh.faces = int32(len(f) / 3)
	copy(mesh.buffer, buf)

	gl.GenVertexArrays(1, &mesh.vao)
	gl.BindVertexArray(mesh.vao)

	gl.GenBuffers(1, &mesh.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.buffer)*4, gl.Ptr(mesh.buffer), gl.STATIC_DRAW)

	mesh.vattrib = uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(mesh.vattrib)
	gl.VertexAttribPointer(mesh.vattrib, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

	mesh.tattrib = uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(mesh.tattrib)
	gl.VertexAttribPointer(mesh.tattrib, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	mesh.nattrib = uint32(gl.GetAttribLocation(program, gl.Str("norm\x00")))
	gl.EnableVertexAttribArray(mesh.nattrib)
	gl.VertexAttribPointer(mesh.nattrib, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(5*4))

	return mesh, nil
}
