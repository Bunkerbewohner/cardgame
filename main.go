// test1 project test1.go
package main

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
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

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
