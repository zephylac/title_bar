package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/zephylac/go-flutter-plugin-statusbar/status_bar"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(&status_bar.StatusBarPlugin{}),
}
