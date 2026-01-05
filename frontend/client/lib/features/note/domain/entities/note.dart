class Note {
  final String? id;
  final String userId;
  final String title;
  final String content;
  final double latitude;
  final double longitude;
  final DateTime? createdAt;

  Note({
    this.id,
    required this.userId,
    required this.title,
    required this.content,
    required this.latitude,
    required this.longitude,
    this.createdAt,
  });
}
