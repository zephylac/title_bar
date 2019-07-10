package status_bar

import (
	"image/color"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

// StatusBarPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/image_picker channel.
type StatusBarPlugin struct {
	BackgroundColor color.RGBA
	Transparent     bool
}

const channelName = "plugins.flutter.io/statusbar"

var _ flutter.Plugin = &StatusBarPlugin{} // compile-time type check

// InitPlugin initializes the path provider plugin.
func (p *StatusBarPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getstatusbarcolor", p.handleGetStatusBar)
	channel.HandleFunc("setstatusbarcolor", p.handleSetStatusBar)
	channel.HandleFunc("setstatusbarwhiteforeground", p.handleSetStatusBarWhiteForeground)
	channel.HandleFunc("getnavigationbarcolor", p.handleGetNavigationBarColor)
	channel.HandleFunc("setnavigationbarwhiteforeground", p.handleSetNavigationBarWhiteForeground)

	return nil
}

func (p *StatusBarPlugin) handleGetStatusBar(arguments interface{}) (reply interface{}, err error) {
	return getStatusBarColor(), nil
}

func (p *StatusBarPlugin) handleSetStatusBar(arguments interface{}) (reply interface{}, err error) {
	color := arguments.(map[interface{}]interface{})["color"].(color.RGBA)

	setStatusBarColor(color)

	return nil, nil
}

func (p *StatusBarPlugin) handleSetStatusBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {

	setStatusBarColor(color.RGBA{255, 255, 255, 255})

	return nil, nil
}

func (p *StatusBarPlugin) handleGetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}

func (p *StatusBarPlugin) handleSetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}

func (p *StatusBarPlugin) handleSetNavigationBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("Desktop doesn't have navigation bar")
}
