package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/zephylac/title_bar"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(&title_bar.TitleBarPlugin{}),
}
