import 'package:flutter/material.dart';
import '../widgets/login_form.dart';
import '../widgets/signup_form.dart';
import 'package:client/core/theme/colors/color_theme.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          bottom: TabBar(
            indicatorColor: context.colors.secondaryColor,
            labelColor: context.colors.secondaryColor,
            unselectedLabelColor: Colors.grey,
            dividerColor: Colors.transparent,
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
