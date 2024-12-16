package main

import (
	"CupidConnector/internal/model"
	"CupidConnector/internal/service"
	"context"
	"fmt"
)

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
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) UpdateConf(config model.Conf) model.Resp {
	err := service.UpdateConf(config)
	if err != nil {
		return model.Resp{Code: model.ResponseCodeError, Msg: err.Error()}
	}
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success"}
}

func (a *App) GetConf() model.Resp {
	config := service.GetConf()
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: config}
}
