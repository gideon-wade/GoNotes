import 'package:flutter/material.dart';
import '../widgets/login_form.dart';
import '../widgets/signup_form.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          bottom: const TabBar(
            indicatorColor: Colors.black,
            labelColor: Colors.black,
            unselectedLabelColor: Colors.grey,
            tabs: [
              Tab(text: 'Login'),
              Tab(text: 'Sign Up'),
            ],
          ),
        ),
        body: const TabBarView(children: [LoginForm(), SignupForm()]),
      ),
    );
  }
}
