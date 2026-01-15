import '../../domain/entities/note.dart';

class NoteModel extends Note {
  NoteModel({
    super.id,
    required super.userId,
    required super.title,
    required super.content,
    required super.latitude,
    required super.longitude,
    super.createdAt,
  });

  factory NoteModel.fromJson(Map<String, dynamic> json) {
    return NoteModel(
      id: json['id'],
      userId: json['userID'],
      title: json['title'],
      content: json['content'],
      latitude: json['latitude'].toDouble(),
      longitude: json['longitude'].toDouble(),
      createdAt: DateTime.parse(json['created_at']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'userID': userId,
      'title': title,
      'content': content,
      'latitude': latitude,
      'longitude': longitude,
    };
  }
}
