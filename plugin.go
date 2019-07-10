package title_bar

import (
	"C"
	"encoding/hex"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "plugins.flutter.io/statusbar"

// TitleBarPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/image_picker channel.
type TitleBarPlugin struct{}

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
	// Retrieve color
	return nil, nil
}

func (p *TitleBarPlugin) handleSetStatusBar(arguments interface{}) (reply interface{}, err error) {
	color := arguments.(map[interface{}]interface{})["color"].(string)
	// animate = false

	colors, _ := hex.DecodeString(color)
	// Retrieve current window
	//C.glfwGetCurrentContext().glClearColor(colors[0], colors[1], colors[2], colors[3])
	//C.glfwGetCurrentContext().glClearColor(glfw.GL_COLOR_BUFFER_BIT)

	return nil, nil
}

func (p *TitleBarPlugin) handleSetStatusBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {
	useWhiteForeground := arguments.(map[interface{}]interface{})["useWhiteForeground"].(bool)
	if useWhiteForeground {
		colors, _ := hex.DecodeString("0000000")
		//C.glfwGetCurrentContext().glClearColor(colors[0], colors[1], colors[2], colors[3])
		//C.glfwGetCurrentContext().glClearColor(glfw.GL_COLOR_BUFFER_BIT)
	}
	return nil, nil
}

func (p *TitleBarPlugin) handleGetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	// Get color
	return nil, nil
}

func (p *TitleBarPlugin) handleSetNavigationBarColor(arguments interface{}) (reply interface{}, err error) {
	color := arguments.(map[interface{}]interface{})["color"].(string)

	colors, _ := hex.DecodeString(color)

	// animate = false
	// Set color & animation
	// 'color': color.value,
	// 'animate': animate,

	//C.glfwGetCurrentContext().glClearColor(colors[0], colors[1], colors[2], colors[3])
	//C.glfwGetCurrentContext().glClearColor(glfw.GL_COLOR_BUFFER_BIT)

	return nil, nil
}

func (p *TitleBarPlugin) handleSetNavigationBarWhiteForeground(arguments interface{}) (reply interface{}, err error) {
	// Set whitforeground navigation bar

	return nil, nil
}

func (p *TitleBarPlugin) useWhiteForeground(backgroundColor string) (reply bool) {
	rgba, _ := hex.DecodeString(backgroundColor)
	// We retrieve luminance from given color
	return 1.05/(rgba[3]+0.05) > 4.5
}
