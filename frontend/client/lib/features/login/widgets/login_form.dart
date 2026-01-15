import 'package:client/core/constants/constants.dart';
import 'package:client/core/utils/utils.dart';
import 'package:client/core/widgets/primary_button.dart';
import 'package:flutter/material.dart';
import 'package:client/features/map/presentation/pages/map_page.dart';

class LoginForm extends StatefulWidget {
  const LoginForm({super.key});

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  final _usernameController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
    _usernameController.dispose();
    _passwordController.dispose();
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
              controller: _usernameController,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(
                context: context,
                hintText: 'Username',
              ),
            ),

            Utils.emptySpace(height: Constants.mSpace),

            TextFormField(
              controller: _passwordController,
              obscureText: true,
              cursorColor: Colors.black,
              decoration: Utils.textFormFieldDecoration(
                context: context,
                hintText: 'Password',
              ),
            ),

            Utils.emptySpace(height: Constants.lSpace),

            PrimaryButton(
              text: 'Login',
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
