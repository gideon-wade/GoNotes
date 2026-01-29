import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class LightColorTheme extends ColorTheme {
  @override
  Color get primaryColor => Color(0xffe9c316);
  @override
  Color get secondaryColor => Color(0xff8eb6d2);
  @override
  Color get accentColor => Color(0xff91be27);
  @override
  Color get surfaceColor => Color(0xfff2e5ba);
  @override
  Color get textColor => Color(0xff1e1101);
  @override
  Color get errorColor => Color(0xFFB00020);
  @override
  Color get onPrimaryColor => Colors.white;
  @override
  Color get onSecondaryColor => Colors.black;
  @override
  Color get onAccentColor => Colors.black;
  @override
  Color get onSurfaceColor => Colors.black;
  @override
  Color get onErrorColor => Colors.white;
  @override
  Brightness get brightness => Brightness.light;
}
