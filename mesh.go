package main

import (
	_ "fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vbo      uint32
	vao      uint32
	vattrib  uint32
	tattrib  uint32
	nattrib  uint32
	faces    int32
	buffer   []float32
	model    mgl32.Mat4
	position mgl32.Vec3
	rotation mgl32.Vec3
	scale    mgl32.Vec3
}

func (m *Mesh) update() {
	ihat := make([]float32, 3)
	jhat := make([]float32, 3)
	khat := make([]float32, 3)
	ihat[0] = 1
	jhat[1] = 1
	khat[2] = 1
	mgl32.vec
	m.model = mgl32.Ident4()
	t := mgl32.Translate3D(m.position.X(), m.position.Y(), m.position.Z())
	rx := mgl32.HomogRotate3D(m.rotation.X(), mgl32.Vec3(ihat))
}

func (m *Mesh) draw(program uint32, modelUniform int32, texture uint32) {
	gl.UseProgram(program)
	gl.UniformMatrix4fv(modelUniform, 1, false, &m.model[0])

	gl.BindVertexArray(m.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.DrawArrays(gl.TRIANGLES, 0, m.faces*2*4)
}
