module github.com/go-flutter-desktop/examples/stocks/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.27.0
	github.com/go-flutter-desktop/plugins/image_picker v0.1.2
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/zephylac/go-flutter-plugin-statusbar/status_bar v0.0.0-00010101000000-000000000000
)

replace github.com/zephylac/go-flutter-plugin-statusbar/status_bar => ../../../go-flutter-plugin-statusbar
