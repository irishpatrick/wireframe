package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
    "github.com/go-gl/gl/v4.1-core/gl"
)

func printSlice(s []float32) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func loadMesh(fn string, program uint32) (Mesh, error) {
    mesh := Mesh{}
    
    fp, err := os.Open(fn)
    check(err)
    defer fp.Close()

    v := make([]float32, 0)
    vt := make([]float32, 0)
    buf := make([]float32, 0)

    // parse
    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        line := scanner.Text()

        if strings.HasPrefix(line, "v ") {
            parts := strings.Split(line, " ")
            for i := 1; i<4; i++ {
                f, err := strconv.ParseFloat(parts[i], 32)
                check(err)
                v = append(v, float32(f))
            }
        }
        if strings.HasPrefix(line, "vt ") {
            parts := strings.Split(line, " ")
            for i := 1; i<3; i++ {
                f, err := strconv.ParseFloat(parts[i], 32)
                check(err)
                vt = append(vt, float32(f))
            }
        }
    }

    // weave
    l := len(v) / 3
    for i := 0; i < l; i++ {
        for j := 0; j < 3; j++ {
            buf = append(buf, v[j])
        }
        for j := 0; j < 2; j++ {
            buf = append(buf, vt[j])
        }
    }

    printSlice(buf)

    copy(mesh.buffer, buf)

	gl.GenVertexArrays(1, &mesh.vao)
    gl.BindVertexArray(mesh.vao)
    
	gl.GenBuffers(1, &mesh.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.buffer)*4, gl.Ptr(mesh.buffer), gl.STATIC_DRAW)

	mesh.vattrib = uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(mesh.vattrib)
	gl.VertexAttribPointer(mesh.vattrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	mesh.tattrib = uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(mesh.tattrib)
	gl.VertexAttribPointer(mesh.tattrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

    return mesh, nil
}
