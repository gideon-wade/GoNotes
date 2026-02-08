import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../cubit/note_cubit.dart';
import '../cubit/note_state.dart';

class CreateNotePage extends StatefulWidget {
  final double latitude;
  final double longitude;

  const CreateNotePage({
    super.key,
    required this.latitude,
    required this.longitude,
  });

  @override
  State<CreateNotePage> createState() => _CreateNotePageState();
}

class _CreateNotePageState extends State<CreateNotePage> {
  final _titleController = TextEditingController();
  final _contentController = TextEditingController();

  void _submit() {
    if (_titleController.text.isEmpty || _contentController.text.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('Need to fill out both title and content'),
        ),
      );
      return;
    }

    context.read<NoteCubit>().createNote(
      userId: "0000004432",
      title: _titleController.text,
      content: _contentController.text,
      latitude: widget.latitude,
      longitude: widget.longitude,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Create Note')),
      body: BlocListener<NoteCubit, NoteState>(
        listener: (context, state) {
          if (state is NoteCreateSuccess) {
            ScaffoldMessenger.of(
              context,
            ).showSnackBar(const SnackBar(content: Text('Note created')));
            Navigator.pop(context);
          } else if (state is NoteError) {
            ScaffoldMessenger.of(
              context,
            ).showSnackBar(SnackBar(content: Text(state.message)));
          }
        },
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                TextField(
                  controller: _titleController,
                  decoration: const InputDecoration(
                    labelText: 'Title',
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 16),
                TextField(
                  controller: _contentController,
                  decoration: const InputDecoration(
                    labelText: 'Content',
                    border: OutlineInputBorder(),
                  ),
                  maxLines: 5,
                ),
                const SizedBox(height: 24),
                BlocBuilder<NoteCubit, NoteState>(
                  builder: (context, state) {
                    return ElevatedButton(
                      onPressed: state is NoteLoading ? null : _submit,
                      child: state is NoteLoading
                          ? const CircularProgressIndicator()
                          : const Text('Create Note'),
                    );
                  },
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  @override
  void dispose() {
    _titleController.dispose();
    _contentController.dispose();
    super.dispose();
  }
}
