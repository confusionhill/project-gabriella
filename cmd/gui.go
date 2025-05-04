package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"com.github/confusionhill/df/private/server/application"
	"com.github/confusionhill/df/private/server/internal/config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/labstack/echo/v4"
)

func startGui(cfg *config.Config, e *echo.Echo, rsc *application.Resources) {
	dfApp := app.New()
	window := dfApp.NewWindow("Server Control")
	hello := widget.NewLabel("Project Gabriella, a DragonFable Private Server")
	startKillBtn := widget.NewButton("Start Server", nil)
	setupBtn := widget.NewButton("Setup Server", nil)
	exitBtn := widget.NewButton("Exit App", nil)
	image := canvas.NewImageFromFile("./public/pages/assets/images/custom-logo.png")
	image.FillMode = canvas.ImageFillOriginal

	// Create a label to hold logs
	logs := widget.NewLabel("")
	logs.Wrapping = fyne.TextWrapWord // Wrap text inside the label

	// Make the log section scrollable
	scrollLogs := container.NewScroll(logs)
	scrollLogs.SetMinSize(fyne.Size{Width: 380, Height: 200}) // Set a min size for the scrollable area

	startKillBtn.OnTapped = func() {
		if serverRunning {
			killServer(cfg, e)
			startKillBtn.SetText("Start Server")
		} else {
			startServer(cfg, e)
			startKillBtn.SetText("Kill Server")
		}
	}

	exitBtn.OnTapped = func() {
		if serverRunning {
			killServer(cfg, e)
		}
		dfApp.Quit()
		log.Println("Exiting app...")
		os.Exit(0)
	}

	setupBtn.OnTapped = func() {
		application.Setup(cfg, rsc)
	}

	content := container.NewVBox(image, hello, startKillBtn, setupBtn, exitBtn, scrollLogs)
	window.SetContent(content)
	logs.SetText(fmt.Sprintf("%s\n%s", logs.Text, "message"))
	// logger.CreateLogger(logs)
	window.ShowAndRun()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	killServer(cfg, e)
}
