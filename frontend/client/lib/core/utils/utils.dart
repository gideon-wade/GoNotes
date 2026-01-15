import 'package:flutter/material.dart';
import '../constants/constants.dart';

class Utils {
  static SizedBox emptySpace({double? width, double? height}) =>
      SizedBox(width: width, height: height);

  static InputDecoration textFormFieldDecoration({
    required BuildContext context,
    String? hintText,
  }) {
    return InputDecoration(
      hintText: hintText,
      filled: true,
      fillColor: Theme.of(context).colorScheme.secondary,
      border: OutlineInputBorder(
        borderRadius: BorderRadius.circular(
          Constants.textFormButtonBorderRadius,
        ),
        borderSide: BorderSide.none,
      ),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(
          Constants.textFormButtonBorderRadius,
        ),
        borderSide: BorderSide.none,
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(
          Constants.textFormButtonBorderRadius,
        ),
        borderSide: BorderSide.none,
      ),
      contentPadding: const EdgeInsets.symmetric(
        horizontal: Constants.buttonContentPadding,
        vertical: Constants.buttonContentPadding,
      ),
    );
  }
}
