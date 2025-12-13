import 'package:flutter_test/flutter_test.dart';
import 'package:client/app.dart';

void main() {
  testWidgets('Check map page loads', (WidgetTester tester) async {
    await tester.pumpWidget(const App());
    expect(find.text('GoNotes Map'), findsOneWidget);
  });
}
