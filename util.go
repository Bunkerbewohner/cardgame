package main

import glfw "github.com/go-gl/glfw3"

// Gets the screen size of the first monitor
func GetScreenSize() (int, int) {
	monitors, err := glfw.GetMonitors()
	if err == nil {
		monitor := monitors[0]
		videoMode, err := monitor.GetVideoMode()
		if err == nil {
			return videoMode.Width, videoMode.Height
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
