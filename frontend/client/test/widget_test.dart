import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:client/app.dart';

void main() {
  testWidgets('Check map page loads', (WidgetTester tester) async {
    await tester.pumpWidget(const MaterialApp(home: App()));

    final loginButton = find.widgetWithText(ElevatedButton, 'Login');
    await tester.tap(loginButton);
    await tester.pumpAndSettle();

    expect(find.text('Profile'), findsOneWidget);
    expect(find.text('Create Note'), findsOneWidget);
    expect(find.text('Notifications'), findsOneWidget);
  });

  testWidgets('Can switch between Login and Sign Up tabs', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(const App());

    final signUpTab = find.ancestor(
      of: find.text('Sign Up'),
      matching: find.byType(Tab),
    );

    await tester.tap(signUpTab);
    await tester.pumpAndSettle();

    expect(find.text('Email'), findsOneWidget);
    expect(find.text('Confirm Password'), findsOneWidget);
  });
}
