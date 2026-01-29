import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class DarkColorTheme extends ColorTheme {
  @override
  Color get primaryColor => Color(0xFFe9c316);
  @override
  Color get secondaryColor => Color(0xFF2d5571);
  @override
  Color get accentColor => Color(0xffabd841);
  @override
  Color get surfaceColor => Color(0xFF231601);
  @override
  Color get textColor => Color(0xFFfef1e1);
  @override
  Color get errorColor => Color(0xFFB00020);
  @override
  Color get onPrimaryColor => Colors.white;
  @override
  Color get onSecondaryColor => Colors.white;
  @override
  Color get onAccentColor => Colors.white;
  @override
  Color get onSurfaceColor => Colors.white;
  @override
  Color get onErrorColor => Colors.white;
  @override
  Brightness get brightness => Brightness.dark;
}
