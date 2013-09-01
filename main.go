// test1 project test1.go
package main

import (
	"fmt"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

func onError(err glfw.ErrorCode, description string) {
	fmt.Printf("%v: %v\n", err, description)
}

func main() {
	fmt.Printf("Launching cardgame...\n")
	glfw.SetErrorCallback(onError)

	if !glfw.Init() {
		panic("OpenGL initialization failed.")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 600, "Cardgame", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetTitle("Cardgame")

	screenWidth, screenHeight := GetScreenSize()
	window.SetPosition(screenWidth/2-400, screenHeight/2-300)

	gl.Init()
	gl.ClearColor(0.5, 0.5, 0.75, 0)
	gl.Ortho(0, 800, 0, 600, 0, 1)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.Color3f(1.0, 0, 0)
		glh.DrawQuadi(100, 100, 200, 300)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
