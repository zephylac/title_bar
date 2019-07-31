package title_bar

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

struct RGBA {
	float red;
	float green;
	float blue;
	float alpha;
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

	rgba.red = [bg redComponent];
	rgba.green = [bg greenComponent];
	rgba.blue = [bg blueComponent];
	rgba.alpha = [bg alphaComponent];

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
		window.styleMask = NSWindowStyleMaskTitled | NSWindowStyleMaskFullScreen;

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

func setTitleBarColor(color color) {
	C.SetTitleBarColor(C.float(float64(color.red)/255), C.float(float64(color.green)/255), C.float(float64(color.blue)/255), C.float(float64(color.alpha)/255))
}

func getTitleBarColor() color {
	temp := C.GetTitleBarColor()

	return color{red: int32(temp.red * 255), green: int32(temp.green * 255), blue: int32(temp.blue * 255), alpha: int32(temp.alpha * 255)}
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
