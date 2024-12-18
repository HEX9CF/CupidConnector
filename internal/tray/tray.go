package tray

import (
	"context"
	_ "embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
			case <-t.ctx.Done():
				return
			case <-quit.ClickedCh:
				t.onExit()
			case <-systray.ClickedCh: // 监听托盘图标左键点击事件
				t.showMainWindow()
			}
		}
	}()
	systray.SetTooltip("Cupid Connector")
}

func (t *Tray) onExit() {
	runtime.Quit(t.ctx)
}

func (t *Tray) showMainWindow() {
	runtime.WindowShow(t.ctx)
}
