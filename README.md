# title_bar

[![pub package](https://img.shields.io/pub/v/title_bar.svg)](https://pub.dartlang.org/packages/title_bar)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://godoc.org/github.com/zephylac/titlebar?status.svg)](http://godoc.org/github.com/zephylac/titlebar)

A Flutter plugin for go-flutter, available for MacOS. It will allow you to customize your title bar by changing the color, name, hiding it, etc.

Note: This plugin is still under development, and some APIs might not be available yet. Feedback welcome and Pull Requests are most welcome!

## Usage

Import as:

```go
import "github.com/zephylac/titlebar"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&title_bar.TitleBarPlugin{}),
```

## More info

If you want to check the current development status go check [here](https://github.com/zephylac/title_bar/projects/1)

## Issues

Please report issues at the [title_bar issue tracker](https://github.com/zephylac/title_bar/issues/).