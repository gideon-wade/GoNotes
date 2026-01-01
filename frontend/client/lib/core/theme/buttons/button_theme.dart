import 'package:flutter/material.dart';
import 'package:client/core/theme/colors/color_theme.dart';
import 'package:client/core/constants/constants.dart';

class AppButtonTheme {
  ColorTheme colorTheme;

  AppButtonTheme(this.colorTheme);

  ElevatedButtonThemeData elevatedButtonThemeData() {
    return ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        backgroundColor: colorTheme.primaryColor,
        foregroundColor: colorTheme.onPrimaryColor,
        minimumSize: const Size(Constants.elevatedButtonWidth, Constants.elevatedButtonHeight),
      ),
    );
  }

  static SizedBox emptySpace({double? width, double? height}) =>
      SizedBox(width: width, height: height);
}