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
      fillColor: Theme.of(context).brightness == Brightness.light
          ? Constants.textFormFieldColorBright
          : Constants.textFormFieldColorDark,
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

  static Widget elevatedTextFormField({
    required BuildContext context,
    String? hintText,
  }) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(
          Constants.textFormButtonBorderRadius,
        ),
        boxShadow: Constants.mediumElevation,
      ),
      child: TextFormField(
        decoration: textFormFieldDecoration(
          context: context,
          hintText: hintText,
        ),
      ),
    );
  }
}
