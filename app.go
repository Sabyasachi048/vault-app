package main

import (
	"context"
	"fmt"

	"github.com/riadafridishibly/wpg/pkg/generator"
	"github.com/riadafridishibly/wpg/pkg/generator/common"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"golang.design/x/clipboard"
)

var lg = logger.NewDefaultLogger()

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := clipboard.Init()
	if err != nil {
		lg.Error("Couldn't initialize clipboard")
	}
}

func (a *App) CopyToClipboard(value string) {
	clipboard.Write(clipboard.FmtText, []byte(value))
}

func (a *App) GenerateNewPassword(seed, password string) (string, error) {
	if seed == "" && password == "" {
		return "", nil
	}

	g := generator.Generator{
		Options: common.Options{
			AllowRepeat:    true,
			IncludeSymbols: true,
			IncludeSpace:   true,
			Length:         23,
		},
	}
	pw, err := g.Generate([]byte(seed), []byte(password))
	if err != nil {
		return "", err
	}

	return pw, nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
