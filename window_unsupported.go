// +build !darwin

package statusbar

import (
	"image/color"

	"github.com/pkg/errors"
)

func setStatusBarColor(color color.RGBA) (error) {
	return errors.New("platform unsupported")
}

func getStatusBarColor() (error) {
	return errors.New("platform unsupported")
}

func setTitleTransparency(titleTransparency bool) (error) {
	return errors.New("platform unsupported")
}
