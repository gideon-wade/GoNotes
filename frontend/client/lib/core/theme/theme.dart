import 'package:client/core/theme/buttons/button_theme.dart';
import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';
import 'package:client/core/theme/colors/light_color_theme.dart';
import 'package:client/core/theme/colors/dark_color_theme.dart';

class AppTheme {
  ColorTheme colorTheme;
  late AppButtonTheme buttonTheme;

  AppTheme({required this.colorTheme}) {
    buttonTheme = AppButtonTheme(colorTheme);
  }

  ThemeData toThemeData() => ThemeData.from(
    useMaterial3: true,
    colorScheme: colorTheme.toColorScheme(),
    textTheme: TextTheme(
      bodyMedium: TextStyle(color: colorTheme.textColor),
      bodyLarge: TextStyle(color: colorTheme.textColor),
    ),
  ).copyWith(elevatedButtonTheme: buttonTheme.elevatedButtonThemeData());

  static AppTheme get light => AppTheme(colorTheme: LightColorTheme());
  static AppTheme get dark => AppTheme(colorTheme: DarkColorTheme());
}
