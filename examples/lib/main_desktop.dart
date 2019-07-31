import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:statusbarapp/main.dart';

void main() {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(new StatusBarApp());
}