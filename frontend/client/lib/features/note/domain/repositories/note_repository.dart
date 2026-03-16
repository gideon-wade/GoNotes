import '../entities/note.dart';

abstract class NoteRepository {
  Future<Note> createNote(Note note);
}