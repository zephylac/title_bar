// +build !darwin

package status_bar

import (
	"image/color"

	"github.com/pkg/errors"
)

func setStatusBarColor(color color.RGBA) {
	return errors.New("platform unsupported")
}

func getStatusBarColor() color.RGBA {
	return errors.New("platform unsupported")
}

func setStatusBarTransparency(transparency bool) {
	return errors.New("platform unsupported")
}

func getStatusBarTransparency() bool {
	return errors.New("platform unsupported")
}

func setStatusBarWidget(hide bool, close bool, minimize bool, resize bool) {
	return errors.New("platform unsupported")
}
