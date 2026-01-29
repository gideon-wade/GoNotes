import 'package:client/features/login/presentation/login_page.dart';
import 'package:flutter/material.dart';
import 'package:client/core/theme/theme.dart';

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'GoNotes',
      theme: AppTheme.dark.toThemeData(),
      darkTheme: AppTheme.dark.toThemeData(),
      themeMode: ThemeMode.system,
      home: const LoginPage(),
    );
  }
}
