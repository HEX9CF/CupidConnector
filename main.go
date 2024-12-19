package main

import (
	"context"
	"cupid-connector/internal/tray"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	// Create an instance of the app structure
	app := NewApp()
	tray.TrayApp = tray.NewTray()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Cupid Connector",
		Width:  600,
		Height: 650,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
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
