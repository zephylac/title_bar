package titlebar

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
SetTitleBarColor(float redValue, float greenValue, float blueValue, float alphaValue) {

	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	NSColor *myColor = [NSColor colorWithCalibratedRed:redValue green:greenValue blue:blueValue alpha:alphaValue];

	window.backgroundColor = myColor;

    return 0;
}

struct RGBA
GetTitleBarColor(void) {

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
SetTitleBarWidget(bool hide, bool close, bool minimize, bool resize) {
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

func setTitleBarColor(color color.RGBA) {
	C.SetTitleBarColor(C.float(float64(color.R)/255), C.float(float64(color.G)/255), C.float(float64(color.B)/255), C.float(float64(color.A)/255))
}

func getTitleBarColor() color.RGBA {
	temp := C.GetTitleBarColor()

	return {"red" : int32(temp.r * 255), "green" : int32(temp.g * 255), "blue" : int32(temp.b * 255), "alpha" : int32(temp.a * 255) }
}

func setTitleBarTransparency(transparency bool) {
	C.SetTitleTransparency(C.bool(transparency))
}

func getTitleBarTransparency() bool {
	return bool(C.GetTitleTransparency())
}

func setTitleBarWidget(hide bool, close bool, minimize bool, resize bool) {
	C.SetTitleBarWidget(C.bool(hide), C.bool(close), C.bool(minimize), C.bool(resize))
}
