package status_bar

import (
	"image/color"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

type StatusBarPlugin struct{}

const channelName = "plugins.flutter.io/statusbar"

var _ flutter.Plugin = &StatusBarPlugin{} // compile-time type check

// InitPlugin initializes the status bar plugin.
func (p *StatusBarPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getstatusbarcolor", p.handleGetStatusBarColor)
	channel.HandleFunc("setstatusbarcolor", p.handleSetStatusBarColor)
	channel.HandleFunc("getstatusbartransparency", p.handleGetStatusBarTransparency)
	channel.HandleFunc("setstatusbartransparency", p.handleSetStatusBarTransparency)
	channel.HandleFunc("hidestatusrar", p.handleHideStatusBar)
	channel.HandleFunc("setstatusbarwidget", p.handleSetStatusBarWidget)

	return nil
}

func (p *StatusBarPlugin) handleGetStatusBarColor(arguments interface{}) (reply interface{}, err error) {
	return getStatusBarColor(), nil
}

func (p *StatusBarPlugin) handleSetStatusBarColor(arguments interface{}) (reply interface{}, err error) {
	color := arguments.(map[interface{}]interface{})["color"].(color.RGBA)

	setStatusBarColor(color)

	return nil, nil
}

func (p *StatusBarPlugin) handleGetStatusBarTransparency(arguments interface{}) (reply interface{}, err error) {
	return getStatusBarTransparency(), nil
}

func (p *StatusBarPlugin) handleSetStatusBarTransparency(arguments interface{}) (reply interface{}, err error) {
	transparency := arguments.(map[interface{}]interface{})["transparency"].(bool)

	setStatusBarTransparency(transparency)

	return nil, nil
}

func (p *StatusBarPlugin) handleHideStatusBar(arguments interface{}) (reply interface{}, err error) {
	hide := arguments.(map[interface{}]interface{})["hide"].(bool)

	hideStatusBar(hide)

	return nil, nil
}

func (p *StatusBarPlugin) handleSetStatusBarWidget(arguments interface{}) (reply interface{}, err error) {
	close := arguments.(map[interface{}]interface{})["close"].(bool)
	minimize := arguments.(map[interface{}]interface{})["minimize"].(bool)

	setStatusBarWidget(close, minimize)

	return nil, nil
}
