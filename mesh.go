package main

import (
    _ "fmt"
    "github.com/go-gl/mathgl/mgl32"

)

type Mesh struct {
    vbo uint32
    vao uint32
    vattrib uint32
    tattrib uint32
    buffer []float32
    position mgl32.Vec3
    rotation mgl32.Vec3
    scale mgl32.Vec3
}

