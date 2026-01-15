import 'package:flutter/material.dart';

abstract class ColorTheme {
  Color get primaryColor;
  Color get secondaryColor;
  Color get accentColor;
  Color get surfaceColor;
  Color get textColor;
  Color get errorColor;
  Color get onPrimaryColor;
  Color get onSecondaryColor;
  Color get onAccentColor;
  Color get onSurfaceColor;
  Color get onErrorColor;
  Brightness get brightness;

  ColorScheme toColorScheme() {
    return ColorScheme(
      primary: primaryColor,
      secondary: secondaryColor,
      tertiary: accentColor,
      surface: surfaceColor,
      error: errorColor,
      onPrimary: onPrimaryColor,
      onSecondary: onSecondaryColor,
      onTertiary: onAccentColor,
      onSurface: onSurfaceColor,
      onError: onErrorColor,
      brightness: brightness,
    );
  }
}
