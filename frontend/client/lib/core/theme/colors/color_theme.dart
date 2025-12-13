import 'package:flutter/material.dart';

abstract class ColorTheme {
  Color get primaryColor;
  Color get secondaryColor;
  Color get surfaceColor;
  Color get errorColor;
  Color get onPrimaryColor;
  Color get onSecondaryColor;
  Color get onSurfaceColor;
  Color get onErrorColor;
  Color get textColor;
  Brightness get brightness;

  ColorScheme toColorScheme() {
    return ColorScheme(
      primary: primaryColor,
      secondary: secondaryColor,
      surface: surfaceColor,
      error: errorColor,
      onPrimary: onPrimaryColor,
      onSecondary: onSecondaryColor,
      onSurface: onSurfaceColor,
      onError: onErrorColor,
      brightness: brightness,
    );
  }
}
