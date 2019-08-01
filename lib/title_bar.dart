/// Call platform code to get/set Title bar or navigation bar
/// background color and foreground brightness.

import 'dart:async';
import 'dart:ui';

import 'package:flutter/services.dart';

class FlutterTitleBar {
  static const MethodChannel _channel = const MethodChannel('plugins.flutter.io/titlebar');

  /// Get the Title bar background color.
  static Future<Color> getTitleBarColor() =>
    _channel.invokeMethod('gettitlebarcolor').then((dynamic response) {
      return response == null ? null : Color.fromARGB(response['alpha'], response['red'], response['green'], response['blue']);
    });

  /// Set the Title bar background color.
  static Future<void> setTitleBarColor(Color color) =>
    _channel.invokeMethod('settitlebarcolor', {
      'alpha': color.alpha,
      'red': color.red,
      'green': color.green,
      'blue': color.blue
    });

  /// Get the Title bar transparency.
  static Future<bool> getTitleBarTransparency() =>
    _channel.invokeMethod('gettitlebartransparency').then((dynamic response) {
      return response == null ? null : response;
    });

  /// Set the Title bar transparency.
  static Future<void> setTitleBarTransparency(bool transparency) =>
    _channel.invokeMethod('settitlebartransparency', {
      'transparency': transparency
    });

  /// Set the Title bar transparency.
  static Future<void> setTitleBarWidget(bool hide, bool close, bool minimize, bool resize) =>
    _channel.invokeMethod('settitlebarwidget', {
      'hide': hide,
      'close': close,
      'minimize': minimize,
      'resize': resize
    });
}

  


  