import 'package:dio/dio.dart';
import 'package:get_it/get_it.dart';
import '../../features/note/data/datasources/note_remote_data_source.dart';
import '../../features/note/data/repositories/note_repository_impl.dart';
import '../../features/note/domain/repositories/note_repository.dart';
import '../../features/note/domain/usecases/create_note.dart';
import '../../features/note/presentation/cubit/note_cubit.dart';

final locator = GetIt.instance;

Future<void> setupLocator() async {
  locator.registerLazySingleton<Dio>(
    () => Dio(BaseOptions(baseUrl: 'http://localhost:8080')),
  );

  locator.registerLazySingleton<NoteRemoteDataSource>(
    () => NoteRemoteDataSource(locator<Dio>()),
  );

  locator.registerLazySingleton<NoteRepository>(
    () => NoteRepositoryImpl(locator<NoteRemoteDataSource>()),
  );

  locator.registerLazySingleton<CreateNote>(
    () => CreateNote(locator<NoteRepository>()),
  );

  locator.registerFactory<NoteCubit>(() => NoteCubit(locator<CreateNote>()));
}
