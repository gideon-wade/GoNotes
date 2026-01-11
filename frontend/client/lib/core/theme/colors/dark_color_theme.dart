import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class DarkColorTheme extends ColorTheme {
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
  Color get primaryColor => Color(0xFFeae12a);
  @override
  Color get secondaryColor => Color(0xFF82db61);
  @override
  Color get surfaceColor => Color(0xFF090801);
  @override
  Color get textColor => Colors.white;
  @override
  Brightness get brightness => Brightness.dark;
}
