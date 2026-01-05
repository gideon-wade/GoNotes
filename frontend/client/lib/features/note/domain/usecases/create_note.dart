import '../entities/note.dart';
import '../repositories/note_repository.dart';

class CreateNote {
  final NoteRepository repository;

  CreateNote(this.repository);

  Future<Note> call({
    required String userId,
    required String title,
    required String content,
    required double latitude,
    required double longitude,
  }) {
    final note = Note(
      userId: userId,
      title: title,
      content: content,
      latitude: latitude,
      longitude: longitude,
    );

    return repository.createNote(note);
  }
}
