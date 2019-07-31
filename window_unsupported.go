// +build !darwin

package titlebar

import (
	"image/color"

	"github.com/pkg/errors"
)

func setTitleBarColor(color color.RGBA) {
	return errors.New("platform unsupported")
}

func getTitleBarColor() color.RGBA {
	return errors.New("platform unsupported")
}

func setTitleBarTransparency(transparency bool) {
	return errors.New("platform unsupported")
}

func getTitleBarTransparency() bool {
	return errors.New("platform unsupported")
}

func setTitleBarWidget(hide bool, close bool, minimize bool, resize bool) {
	return errors.New("platform unsupported")
}
