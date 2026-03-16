import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class LightColorTheme extends ColorTheme {
  @override
  Color get errorColor => Color(0xFFB00020);
  @override
  Color get onErrorColor => Colors.white;
  @override
  Color get onPrimaryColor => Colors.white;
  @override
  Color get onSecondaryColor => Colors.black;
  @override
  Color get onSurfaceColor => Colors.black;
  @override
  Color get primaryColor => Color(0xffcdc415);
  @override
  Color get secondaryColor => Color(0xff67af4d);
  @override
  Color get surfaceColor => Colors.white;
  @override
  Color get textColor => Colors.black;
  @override
  Brightness get brightness => Brightness.light;
}
