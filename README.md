# title_bar

A Flutter plugin for go-flutter, available for MacOS. It will allow you to customize your title bar by changing the color, name, hiding it, etc.

Note: This plugin is still under development, and some APIs might not be available yet. Feedback welcome and Pull Requests are most welcome!

## Usage

Import as:

```go
import "github.com/zephylac/title_bar"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&title_bar.TitleBarPlugin{}),
```

## Issues

Please report issues at the [title_bar issue tracker](https://github.com/zephylac/title_bar/issues/).