import '../../domain/entities/note.dart';
import '../../domain/repositories/note_repository.dart';
import '../datasources/note_remote_data_source.dart';
import '../models/note_model.dart';

class NoteRepositoryImpl implements NoteRepository {
  final NoteRemoteDataSource remoteDataSource;

  NoteRepositoryImpl(this.remoteDataSource);

  @override
  Future<Note> createNote(Note note) async {
    final noteModel = NoteModel(
      title: note.title,
      content: note.content,
      latitude: note.latitude,
      longitude: note.longitude,
    );

    return await remoteDataSource.createNote(noteModel);
  }
}