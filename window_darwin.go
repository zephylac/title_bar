package title_bar

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

struct RGBA {
	float r;
	float g;
	float b;
	float a;
};

int
SetStatusBarColor(float redValue, float greenValue, float blueValue, float alphaValue) {

	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	NSColor *myColor = [NSColor colorWithCalibratedRed:redValue green:greenValue blue:blueValue alpha:alphaValue];

	window.backgroundColor = myColor;

    return 0;
}

struct RGBA
GetStatusBarColor(void) {

	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	NSColor *bg =  [window valueForKey:@"backgroundColor"];

	struct RGBA rgba;

	rgba.r = [bg redComponent];
	rgba.g = [bg greenComponent];
	rgba.b = [bg blueComponent];
	rgba.a = [bg alphaComponent];

	return rgba;

}

bool
GetTitleTransparency() {

	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	return [window valueForKey:@"titlebarAppearsTransparent"];
}

int
SetTitleTransparency(bool titleTransparency) {

	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	window.titlebarAppearsTransparent =  titleTransparency;

    return 0;
}

int
SetStatusBarWidget(bool hide, bool close, bool minimize, bool resize) {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	if (hide) {
		window.styleMask = NSWindowStyleMaskBorderless;
	} else {
		window.styleMask = NSWindowStyleMaskTitled | NSFullSizeContentViewWindowMask;

		if (close) {
			window.styleMask |= NSWindowStyleMaskClosable;
		} else {
			window.styleMask &= ~NSWindowStyleMaskClosable;
		}

		if (minimize) {
			window.styleMask |= NSWindowStyleMaskMiniaturizable;
		} else {
			window.styleMask &= ~NSWindowStyleMaskMiniaturizable;
		}

		if (resize) {
			window.styleMask |= NSWindowStyleMaskResizable;
		} else {
			window.styleMask &= ~NSWindowStyleMaskResizable;
		}
	}

	return 0;
}
*/
import "C"

import "image/color"

func setStatusBarColor(color color.RGBA) {
	C.SetStatusBarColor(C.float(float64(color.R)/255), C.float(float64(color.G)/255), C.float(float64(color.B)/255), C.float(float64(color.A)/255))
}

func getStatusBarColor() color.RGBA {
	temp := C.GetStatusBarColor()

	return color.RGBA{uint8(temp.r * 255), uint8(temp.g * 255), uint8(temp.b * 255), uint8(temp.a * 255)}
}

func setStatusBarTransparency(transparency bool) {
	C.SetTitleTransparency(C.bool(transparency))
}

func getStatusBarTransparency() bool {
	return bool(C.GetTitleTransparency())
}

func setStatusBarWidget(hide bool, close bool, minimize bool, resize bool) {
	C.SetStatusBarWidget(C.bool(hide), C.bool(close), C.bool(minimize), C.bool(resize))
}
