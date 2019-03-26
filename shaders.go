package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/go-gl/gl/v4.1-core/gl"
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

	buffer := make([]byte, stat.Size())
	_, err = f.Read(buffer)
	check(err)

	f.Close()

	return string(buffer), nil
}

func compileShader(source string, shader_type uint32) (uint32, error) {
	shader := gl.CreateShader(shader_type)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newProgram(vertex_fn string, fragment_fn string) (uint32, error) {
	vertex_src, err := readFile(vertex_fn)
	check(err)
	fragment_src, err := readFile(fragment_fn)
	check(err)

	vsid, err := compileShader(vertex_src, gl.VERTEX_SHADER)
	check(err)
	fsid, err := compileShader(fragment_src, gl.FRAGMENT_SHADER)
	check(err)

	pid := gl.CreateProgram()
	gl.AttachShader(pid, vsid)
	gl.AttachShader(pid, fsid)

	gl.LinkProgram(pid)

	var status int32
	gl.GetProgramiv(pid, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(pid, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(pid, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link pid: %v", log)
	}

	gl.DeleteShader(vsid)
	gl.DeleteShader(fsid)

	return pid, nil
}