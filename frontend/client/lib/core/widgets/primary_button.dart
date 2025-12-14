import 'package:flutter/material.dart';
import 'package:client/core/constants/constants.dart';

class PrimaryButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;

  const PrimaryButton({super.key, required this.text, required this.onPressed});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: Constants.elevatedButtonWidth,
      height: Constants.elevatedButtonHeight,
      child: ElevatedButton(
        onPressed: onPressed,
        style: ElevatedButton.styleFrom(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(
              Constants.primaryButtonBorderRadius,
            ),
          ),
        ),
        child: Text(text),
      ),
    );
  }
}
