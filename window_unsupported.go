// +build !darwin

package title_bar

import (
	"github.com/pkg/errors"
)

func setTitleBarColor(color color) {
	return errors.New("platform unsupported")
}

func getTitleBarColor() color {
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
