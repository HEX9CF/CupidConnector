package tray

import (
	"context"
	"cupid-connector/internal/data"
	_ "embed"
	"fmt"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var TrayApp *Tray

//go:embed icon.ico
var icon []byte

type Tray struct {
	ctx context.Context
}

func NewTray() *Tray {
	return &Tray{}
}

func (t *Tray) Startup(ctx context.Context) {
	t.ctx = ctx
	t.Run()
}

func (t *Tray) Run() {
	systray.Run(t.onReady, t.onExit)
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)

	show := systray.AddMenuItem("主窗口", "")
	quit := systray.AddMenuItem("退出", "")
	go func() {
		for {
			select {
			case <-show.ClickedCh:
				t.showMainWindow()
			case <-quit.ClickedCh:
				t.onExit()
			case <-t.ctx.Done():
				return
			}
		}
	}()
	t.TooltipRefresh()
}

func (t *Tray) TooltipRefresh() {
	systray.SetTooltip(fmt.Sprintf("剩余流量：%.3f MB", data.Info.Overall-data.Info.Used))
}

func (t *Tray) onExit() {
	runtime.Quit(t.ctx)
}

func (t *Tray) showMainWindow() {
	runtime.WindowShow(t.ctx)
}

func (t *Tray) IsStartUp() bool {
	return t.ctx != nil
}
