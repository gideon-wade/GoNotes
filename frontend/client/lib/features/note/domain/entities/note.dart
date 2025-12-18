class Note {
  final String? id;
  final String title;
  final String content;
  final double latitude;
  final double longitude;
  final DateTime? createdAt;

  Note({
    this.id,
    required this.title,
    required this.content,
    required this.latitude,
    required this.longitude,
    this.createdAt,
  });
}