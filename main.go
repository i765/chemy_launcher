package main

import (
	"embed"
	"fmt"
	"log"
	"os/exec"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.bug.st/serial"
)

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

func (a *App) ShowMessage(message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Информация",
		Message: message,
	})
}

func (a *App) ShowError(title string, mess string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   title,
		Message: fmt.Sprintf("%s", mess),
	})
}

// Запуск с ожиданием результата и получением вывода
func (a *App) RunCmd(command string, args ...string) (bool, error) {
	cmd := exec.Command(command, args...)

	// Полностью игнорируем stdout и stderr, чтобы не тратить память
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Run() сам запускает процесс и ждет его окончания
	err := cmd.Run()
	if err != nil {
		// Если программа вернула ошибку (код не 0) или не смогла запуститься
		return false, fmt.Errorf("программа завершилась с ошибкой: %w", err)
	}
	// Программа успешно выполнилась и закрылась
	return true, nil
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Запуск программы",
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
