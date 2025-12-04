import Flutter
import UIKit

class Env {
  static String get IOS_API_KEY => dotenv.env['IOS_API_KEY']!;
}

@main
@objc class AppDelegate: FlutterAppDelegate {
  override func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?
  ) -> Bool {
    GeneratedPluginRegistrant.register(with: self)
    GMSServices.provideAPIKey(Env.IOS_API_KEY)
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}
