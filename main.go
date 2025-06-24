package main

import (
	"context"
	"cupid-connector/backend/application"
	"cupid-connector/backend/tray"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	// Create an instance of the app structure
	app := application.NewApp()
	tray.TrayApp = tray.NewTray()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Cupid Connector",
		Width:  1050,
		Height: 650,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:         logger.ERROR,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)

			// 启动托盘（必须在最后启动）
			tray.TrayApp.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			tray.TrayApp,
		},
		DisableResize: true,
		Frameless:     true,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
