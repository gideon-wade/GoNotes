import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class DarkColorTheme extends ColorTheme {
  @override
  Color get primaryColor => Color(0xFFe9c316);
  @override
  Color get secondaryColor => Color(0xFF33305a);
  @override
  Color get accentColor => Color(0xff403d7b);
  @override
  Color get surfaceColor => Color(0xFF1d160c);
  @override
  Color get textColor => Color(0xFF7a75db);
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
  Brightness get brightness => Brightness.dark;
}
