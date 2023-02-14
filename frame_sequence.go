package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

const (
	frameWidth  = 640
	frameHeight = 480
)

func loadView() {
	myApp := app.New()
	myWindow := myApp.NewWindow("--------oo00oo-----oo00oo----------")
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.CenterOnScreen()
	myWindow.SetContent(widget.NewButton("Run", func() {
		run(myApp)
	}))
	myWindow.ShowAndRun()
}

func run(myApp fyne.App) {

	// create a new window with a fixed size
	myWindow := myApp.NewWindow("-----------oo00oo-----oo00oo----------")
	myWindow.Resize(fyne.NewSize(frameWidth, frameHeight))

	// create a new container to hold the image views
	//imageContainer := container.New(layout.NewMaxLayout())

	// load the image frames from disk
	framePaths, err := filepath.Glob("frames/heart_animation_*.png")
	if err != nil {
		fmt.Println("Error: could not find image frames:", err)
		return
	}

	// sort the frame paths by name
	sort.Slice(framePaths, func(i, j int) bool {

		// extract the integer value from the file name using strconv.Atoi
		iVal, _ := strconv.Atoi(framePaths[i][len("frames/heart_animation_") : len(framePaths[i])-len(".png")])
		jVal, _ := strconv.Atoi(framePaths[j][len("frames/heart_animation_") : len(framePaths[j])-len(".png")])

		return iVal < jVal
	})

	// create an image view
	imageView := canvas.NewImageFromFile("frames/heart_animation_0.png")
	imageView.FillMode = canvas.ImageFillOriginal
	imageView.SetMinSize(fyne.NewSize(frameWidth, frameHeight))
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewVBox(
		imageView,
		widget.NewButton("Close", func() {
			myApp.Quit()
			os.Exit(1)
		}),
	))

	// preload the images into memory
	var imageResources []fyne.Resource
	for i := 0; i < len(framePaths); i++ {
		framePath := framePaths[i]

		// load the frame image file as a resource
		imageResource, err := fyne.LoadResourceFromPath(framePath)
		if err != nil {
			fmt.Println("Error: could not load image resource:", err)
			return
		}
		imageResources = append(imageResources, imageResource)
	}

	frameIndex := 0
	go func() {
		for {
			imageView.Resource = imageResources[frameIndex]
			frameIndex = (frameIndex + 1) % len(imageResources)
			time.Sleep(67 * time.Millisecond) // adjust to desired frame rate
			imageView.Refresh()

		}
	}()
	// show the window and start the main loop
	myWindow.Show()
}
