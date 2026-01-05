import '../../domain/entities/note.dart';

abstract class NoteState {}

class NoteInitial extends NoteState {}

class NoteLoading extends NoteState {}

class NoteCreateSuccess extends NoteState {
  final Note note;
  NoteCreateSuccess(this.note);
}

class NoteError extends NoteState {
  final String message;
  NoteError(this.message);
}