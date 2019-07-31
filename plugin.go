package title_bar

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// TitleBarPlugin handles method calls to
// the plugins.flutter.io/titlebar channel.
type TitleBarPlugin struct{}

const channelName = "plugins.flutter.io/titlebar"

var _ flutter.Plugin = &TitleBarPlugin{} // compile-time type check

// Color struct
type color struct {
	alpha int32
	red   int32
	green int32
	blue  int32
}

// InitPlugin initializes the title bar plugin.
func (p *TitleBarPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("gettitlebarcolor", p.handleGetTitleBarColor)
	channel.HandleFunc("settitlebarcolor", p.handleSetTitleBarColor)
	channel.HandleFunc("gettitlebartransparency", p.handleGetTitleBarTransparency)
	channel.HandleFunc("settitlebartransparency", p.handleSetTitleBarTransparency)
	channel.HandleFunc("settitlebarwidget", p.handleSetTitleBarWidget)

	setTitleBarTransparency(true)

	return nil
}

func (p *TitleBarPlugin) handleGetTitleBarColor(arguments interface{}) (reply interface{}, err error) {
	return getTitleBarColor(), nil
}

func (p *TitleBarPlugin) handleSetTitleBarColor(arguments interface{}) (reply interface{}, err error) {
	red := arguments.(map[interface{}]interface{})["red"].(int32)
	green := arguments.(map[interface{}]interface{})["green"].(int32)
	blue := arguments.(map[interface{}]interface{})["blue"].(int32)
	alpha := arguments.(map[interface{}]interface{})["alpha"].(int32)

	setTitleBarColor(color{red: int32(red), green: int32(green), blue: (blue), alpha: (alpha)})

	return nil, nil
}

func (p *TitleBarPlugin) handleGetTitleBarTransparency(arguments interface{}) (reply interface{}, err error) {
	return getTitleBarTransparency(), nil
}

func (p *TitleBarPlugin) handleSetTitleBarTransparency(arguments interface{}) (reply interface{}, err error) {
	transparency := arguments.(map[interface{}]interface{})["transparency"].(bool)

	setTitleBarTransparency(transparency)

	return nil, nil
}

func (p *TitleBarPlugin) handleSetTitleBarWidget(arguments interface{}) (reply interface{}, err error) {
	hide := arguments.(map[interface{}]interface{})["hide"].(bool)
	close := arguments.(map[interface{}]interface{})["close"].(bool)
	minimize := arguments.(map[interface{}]interface{})["minimize"].(bool)
	resize := arguments.(map[interface{}]interface{})["resize"].(bool)

	setTitleBarWidget(hide, close, minimize, resize)

	return nil, nil
}
