# title_bar

[![pub package](https://img.shields.io/pub/v/title_bar.svg)](https://pub.dartlang.org/packages/title_bar)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://godoc.org/github.com/zephylac/title_bar?status.svg)](http://godoc.org/github.com/zephylac/title_bar)

A Flutter plugin for go-flutter, available for MacOS. It will allow you to customize your title bar by changing the color, name, hiding it, etc.

Note: This plugin is still under development, and some APIs might not be available yet. Feedback welcome and Pull Requests are most welcome!

## Getting started

### Flutter side

In the pubspec.yaml of your flutter project, add the following dependency:

```yml
dependencies:
  ...
  flutter_slidable: "^0.2.0"
```

In your library add the following import:

```dart
import 'package:title_bar/title_bar.dart';
```

For help getting started with Flutter, view the online [documentation](https://flutter.dev/).

### go-flutter side

Import as:

```go
import "github.com/zephylac/title_bar"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&title_bar.TitleBarPlugin{}),
```

## Changelog

Please see the [changelog](https://github.com/zephylac/title_bar/blob/master/CHANGELOG.md) page to know what's recently changed.

## More info

If you want to check the current development status go check [here](https://github.com/zephylac/title_bar/projects/1)

## Contributions

Feel free to contribute to this project.

If you find a bug or want a feature, but don't know how to fix/implement it, please fill an [issue](https://github.com/zephylac/title_bar/issues/).
If you fixed a bug or implemented a feature, please send a pull [request](https://github.com/zephylac/title_bar/pulls/).
