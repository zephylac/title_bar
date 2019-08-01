import 'package:flutter/material.dart';
import 'package:title_bar/title_bar.dart';

void main() {
  runApp(new TitleBarApp());
}
class TitleBarApp extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: 'Generated App',
      theme: new ThemeData(
        primarySwatch: Colors.blue,
        primaryColor: const Color(0xFF2196f3),
        accentColor: const Color(0xFF2196f3),
        canvasColor: const Color(0xFFfafafa),
      ),
      home: new MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key}) : super(key: key);
  @override
  _MyHomePageState createState() => new _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
    int _red = 255;
    int _blue = 255;
    int _green = 255;
    int _alpha = 255;

    bool _hide = false;
    bool _close = true;
    bool _minimize = true;
    bool _resize = true;

    @override
    Widget build(BuildContext context) {
      return new Scaffold(
        appBar: new AppBar(
          title: new Text('Example title bar'),
          ),
        body:
          new Column(
            mainAxisAlignment: MainAxisAlignment.start,
            mainAxisSize: MainAxisSize.max,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              new Text(
                "Red: $_red",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFFff0000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Slider(key: null, onChanged: sliderRedChanged, min: 0, max: 255 ,value:_red.toDouble(),),
    
              new Text(
              "Blue: $_blue",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF001aff),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Slider(key: null, onChanged: sliderBlueChanged, min: 0, max: 255 ,value:_blue.toDouble(),),
    
              new Text(
              "Green: $_green",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF4dff00),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Slider(key: null, onChanged: sliderGreenChanged, min: 0, max: 255 ,value:_green.toDouble(),),

              new Text(
              "Alpha $_alpha",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF000000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Slider(key: null, onChanged: sliderAlphaChanged, min: 0, max: 255 ,value:_alpha.toDouble(),),

              new Text(
              "Hide title bar",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF000000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Switch(onChanged: switchHideTitleBarChanged, value:_hide),
    
              new Text(
              "Hide close",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF000000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Switch(onChanged: switchHideCloseChanged, value:_close),
    
              new Text(
              "Hide minimize",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF000000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Switch(onChanged: switchHideMinimizeChanged, value:_minimize),

              new Text(
              "Hide resize",
                style: new TextStyle(fontSize:12.0,
                color: const Color(0xFF000000),
                fontWeight: FontWeight.w200,
                fontFamily: "Roboto"),
              ),
    
              new Switch(onChanged: switchHideResizeChanged, value:_resize)
            ]
    
          ),
    
      );
    }
    
    void sliderRedChanged(double value) async {
      setState(() => _red = value.round());
      FlutterTitleBar.setTitleBarColor(Color.fromARGB(_alpha, _red, _green, _blue));
    }

    void sliderBlueChanged(double value) async {
      setState(() => _blue = value.round());
      FlutterTitleBar.setTitleBarColor(Color.fromARGB(_alpha, _red, _green, _blue));
    }

    void sliderGreenChanged(double value) async {
      setState(() => _green = value.round());
      FlutterTitleBar.setTitleBarColor(Color.fromARGB(_alpha, _red, _green, _blue));
    }
    
    void sliderAlphaChanged(double value) async {
      setState(() => _alpha = value.round());
      FlutterTitleBar.setTitleBarColor(Color.fromARGB(_alpha, _red, _green, _blue));
    }

    void switchHideTitleBarChanged(bool value) async {
      setState(() => _hide = value);
      FlutterTitleBar.setTitleBarWidget(_hide, _close, _minimize, _resize);
    }

    void switchHideCloseChanged(bool value) async {
      setState(() => _close = value);
      FlutterTitleBar.setTitleBarWidget(_hide, _close, _minimize, _resize);
    }

    void switchHideMinimizeChanged(bool value) async {
      setState(() => _minimize = value);
      FlutterTitleBar.setTitleBarWidget(_hide, _close, _minimize, _resize);
    }
    
    void switchHideResizeChanged(bool value) async {
      setState(() => _resize = value);
      FlutterTitleBar.setTitleBarWidget(_hide, _close, _minimize, _resize);
    }
}