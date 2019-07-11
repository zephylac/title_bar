package status_bar

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

	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	NSColor *myColor = [NSColor colorWithCalibratedRed:redValue green:greenValue blue:blueValue alpha:alphaValue];
	[window setBackgroundColor: myColor];

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

	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	[window setTitlebarAppearsTransparent: titleTransparency];

    return 0;
}


int
HideStatusBar(bool hideStatus) {
	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	if(hideStatus) {
		[window setStyleMask:NSWindowStyleMaskBorderless];
	}

	return 0;
}

int
SetStatusBarWidget(bool close, bool minimize) {
	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	if(close && minimize) {
		[window setStyleMask:NSWindowStyleMaskTitled | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskClosable];
	} else if(close) {
		[window setStyleMask:NSWindowStyleMaskTitled | NSWindowStyleMaskClosable];
	} else if(minimize) {
		[window setStyleMask:NSWindowStyleMaskTitled | NSWindowStyleMaskMiniaturizable];
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

func hideStatusBar(hideStatusbar bool) {
	C.HideStatusBar(C.bool(hideStatusbar))
}

func setStatusBarWidget(close bool, minimize bool) {
	C.SetStatusBarWidget(C.bool(close), C.bool(minimize))
}