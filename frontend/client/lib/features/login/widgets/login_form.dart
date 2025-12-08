import 'package:flutter/material.dart';
import 'package:client/features/map/presentation/pages/map_page.dart';

class LoginForm extends StatelessWidget {
  const LoginForm({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        TextFormField(decoration: const InputDecoration(labelText: 'Username')),
        TextFormField(
          decoration: const InputDecoration(labelText: 'Password'),
          obscureText: true,
        ),
        ElevatedButton(
          onPressed: () {
            Navigator.pushReplacement(
              context,
              MaterialPageRoute(builder: (_) => const MapPage()),
            );
          },
          child: const Text('Login'),
        ),
      ],
    );
  }
}
