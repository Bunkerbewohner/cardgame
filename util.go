package main

import glfw "github.com/go-gl/glfw3"
import "math"

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

func Round(num float64) float64 {
	if num >= 0 {
		return math.Floor(num + 0.5)
	} else {
		return math.Ceil(num - 0.5)
	}
}
