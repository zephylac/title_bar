package statusbar

import (
	"image/color"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

// TitleBarPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/image_picker channel.
type TitleBarPlugin struct {
	BackgroundColor color.RGBA
	Transparent     bool
}

const channelName = "plugins.flutter.io/statusbar"

var _ flutter.Plugin = &TitleBarPlugin{} // compile-time type check

// InitPlugin initializes the path provider plugin.
func (p *TitleBarPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getstatusbarcolor", p.handleGetStatusBar)
	channel.HandleFunc("setstatusbarcolor", p.handleSetStatusBar)
	channel.HandleFunc("setstatusbarwhiteforeground", p.handleSetStatusBarWhiteForeground)
	channel.HandleFunc("getnavigationbarcolor", p.handleGetNavigationBarColor)
	channel.HandleFunc("setnavigationbarwhiteforeground", p.handleSetNavigationBarWhiteForeground)

	return nil
}

func (p *TitleBarPlugin) handleGetStatusBar(arguments interface{}) (reply interface{}, err error) {
	return getStatusBarColor(), nil
}

func (p *TitleBarPlugin) handleSetStatusBar(arguments interface{}) (reply interface{}, err error) {
	color := arguments.(map[interface{}]interface{})["color"].(color.RGBA)

	setStatusBarColor(color)

	return nil, nil
}

func (p *TitleBarPlugin) handleSetStatusBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {

	setStatusBarColor(color.RGBA{255, 255, 255, 255})

	return nil, nil
}

func (p *TitleBarPlugin) handleGetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}

func (p *TitleBarPlugin) handleSetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}

func (p *TitleBarPlugin) handleSetNavigationBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}
