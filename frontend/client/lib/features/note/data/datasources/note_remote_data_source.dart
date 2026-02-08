import 'package:dio/dio.dart';
import '../models/note_model.dart';

class NoteRemoteDataSource {
  final Dio dio;

  NoteRemoteDataSource(this.dio);

  Future<NoteModel> createNote(NoteModel note) async {
    final response = await dio.post('/api/v1/notes', data: note.toJson());

    return NoteModel.fromJson(response.data);
  }
}
