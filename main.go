package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.bug.st/serial"
)

//go:embed all:frontend/dist
var assets embed.FS

// GetAvailablePorts запрашивает порты и возвращает их для фронтенда
func (a *App) GetAvailablePorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Printf("Ошибка получения портов: %v", err)
		return []string{} // возвращаем пустой массив в случае ошибки
	}
	return ports
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Launcher",
		Width:         550,
		Height:        300,
		Frameless:     true,
		AlwaysOnTop:   true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 32, G: 32, B: 32, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
