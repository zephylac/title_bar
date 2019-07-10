package statusbar

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
int
ChangeColor(int redValue, int greenValue, int blueValue, int alphaValue) {

	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	NSColor *myColor = [NSColor colorWithCalibratedRed:redValue green:greenValue blue:blueValue alpha:alphaValue];
	[window setBackgroundColor: myColor];

    return 0;
}
*/
import "C"

import flutter "github.com/go-flutter-desktop/go-flutter"
import "github.com/go-flutter-desktop/go-flutter/plugin"

// TitleBarPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/image_picker channel.
type TitleBarPlugin struct{}

var _ flutter.Plugin = &TitleBarPlugin{} // compile-time type check

// InitPlugin initializes the path provider plugin.
func (p *TitleBarPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	C.ChangeColor(1, 0, 0, 1)

	return nil
}

/* func (p *TitleBarPlugin) useWhiteForeground(backgroundColor string) (reply bool) {
	rgba, _ := hex.DecodeString(backgroundColor)
	// We retrieve luminance from given color
	return 1.05/(rgba[3]+0.05) > 4.5
} */
