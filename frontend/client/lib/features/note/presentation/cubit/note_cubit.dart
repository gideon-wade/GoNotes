import 'package:flutter_bloc/flutter_bloc.dart';
import '../../domain/usecases/create_note.dart';
import 'note_state.dart';

class NoteCubit extends Cubit<NoteState> {
  final CreateNote createNoteUseCase;

  NoteCubit(this.createNoteUseCase) : super(NoteInitial());

  Future<void> createNote({
    required String title,
    required String content,
    required double latitude,
    required double longitude,
  }) async {
    emit(NoteLoading());
    try {
      final note = await createNoteUseCase(
        title: title,
        content: content,
        latitude: latitude,
        longitude: longitude,
      );
      emit(NoteCreateSuccess(note));
    } catch (e) {
      emit(NoteError(e.toString()));
    }
  }
}