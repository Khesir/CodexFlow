package cmd

import (
	server "app/internal"
	"context"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

var embeddedFrontend embed.FS

type App struct{}

func NewApp() *App { return &App{} }

func (a *App) Startup(ctx context.Context) {
	// optional: do startup logic
}

func Main() {
	// Start Gin server in background; provide embedded frontend FS so the
	// same API + static serve works in the desktop packaged app.
	go func() {
		// pass the raw embed FS; server.Start will take care of fs.Sub internally
		if err := server.Start(":8080", embeddedFrontend); err != nil {
			log.Println("server stopped:", err)
		}
	}()

	app := NewApp()
	err := wails.Run(&options.App{
		Title:     "CodexFlow",
		Width:     1200,
		Height:    800,
		Assets:    embeddedFrontend,
		OnStartup: app.Startup,
	})
	if err != nil {
		log.Fatal("wails run error:", err)
	}
}
