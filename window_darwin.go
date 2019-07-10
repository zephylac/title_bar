package statusbar

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

int
SetTitleTransparency(bool titleTransparency) {

	id window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];

	[window setTitlebarAppearsTransparent: titleTransparency];

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

func setTitleTransparency(titleTransparency bool) {
	C.SetTitleTransparency(C.bool(titleTransparency))
}
