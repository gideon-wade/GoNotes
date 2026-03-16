import 'package:client/core/constants/constants.dart';
import 'package:client/core/utils/utils.dart';
import 'package:client/core/widgets/primary_button.dart';
import 'package:flutter/material.dart';
import '../../map/presentation/map_page.dart';

class SignupForm extends StatefulWidget {
  const SignupForm({super.key});

  @override
  State<SignupForm> createState() => _SignupFormState();
}

class _SignupFormState extends State<SignupForm> {
  final _emailController = TextEditingController();
  final _usernameController = TextEditingController();  
  final _passwordController = TextEditingController();
  final _confirmPasswordController = TextEditingController();

  @override
  void dispose() {
    _emailController.dispose();
    _usernameController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Padding(
        padding: const EdgeInsets.all(Constants.lSpace),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextFormField(
              controller: _emailController,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(hintText: 'Email'),
            ),

            Utils.emptySpace(height: Constants.mSpace),

            TextFormField(
              controller: _usernameController,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(hintText: 'Username'),
            ),

            Utils.emptySpace(height: Constants.mSpace),

            TextFormField(
              controller: _passwordController,
              obscureText: true,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(hintText: 'Password'),
            ),

            Utils.emptySpace(height: Constants.mSpace),

            TextFormField(
              controller: _confirmPasswordController,
              obscureText: true,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(
                hintText: 'Confirm Password',
              ),
            ),

            Utils.emptySpace(height: Constants.lSpace),

            PrimaryButton(
              text: 'Sign Up',
              onPressed: () {
                Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(builder: (_) => const MapPage()),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}
