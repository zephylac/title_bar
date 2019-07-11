// +build !darwin

package status_bar

import (
	"image/color"

	"github.com/pkg/errors"
)

func setStatusBarColor(color color.RGBA) error {
	return errors.New("platform unsupported")
}

func getStatusBarColor() error {
	return errors.New("platform unsupported")
}

func setStatusBarTransparency(transparency bool) error {
	return errors.New("platform unsupported")
}

func getStatusBarTransparency() bool {
	return errors.New("platform unsupported")
}

func hideStatusBar(hideTitle bool) {
	return errors.New("platform unsupported")
}

func setStatusBarWidget(close bool, minimize bool) {
	return errors.New("platform unsupported")
}
