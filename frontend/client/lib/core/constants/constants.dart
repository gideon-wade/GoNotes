import 'package:flutter/material.dart';

class Constants {
  static const String appName = 'GoNotes';
  // Spacing
  static const double xsSpace = 4.0;
  static const double sSpace = 8.0;
  static const double mSpace = 16.0;
  static const double lSpace = 24.0;
  static const double xlSpace = 32.0;
  // Buttons
  static const double elevatedButtonHeight = 48.0;
  static const double elevatedButtonWidth = 144.0;
  static const double textFormButtonBorderRadius = 12.0;
  static const double buttonContentPadding = 16.0;
  static const double primaryButtonBorderRadius = 30.0;
  static const double primaryButtonFontSize = 18.0;

  // Colors
  static final Color? greyTransparentColor = Colors.grey[200];
  static final Color textFormFieldColorBright = Color(0XFFf3ece2);
  static final Color textFormFieldColorDark = Color(0XFF44370d);

  // Elevations
  static List<BoxShadow> get subtleElevation => [
    BoxShadow(
      color: Colors.black.withAlpha(20),
      offset: const Offset(0, 2),
      blurRadius: 4,
      spreadRadius: 0,
    ),
    BoxShadow(
      color: Colors.black.withAlpha(10),
      offset: const Offset(0, 1),
      blurRadius: 2,
      spreadRadius: 0,
    ),
  ];

  static List<BoxShadow> get mediumElevation => [
    BoxShadow(
      color: Colors.black.withAlpha(25),
      offset: const Offset(0, 4),
      blurRadius: 8,
      spreadRadius: 0,
    ),
    BoxShadow(
      color: Colors.black.withAlpha(15),
      offset: const Offset(0, 2),
      blurRadius: 4,
      spreadRadius: 0,
    ),
  ];
}
