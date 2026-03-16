import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:geolocator/geolocator.dart';
import 'package:client/core/di/injection_container.dart';
import 'package:client/features/note/presentation/cubit/note_cubit.dart';
import 'package:client/features/note/presentation/pages/create_note_page.dart';

class MapBottomBar extends StatefulWidget {
  final Position? currentPosition;

  const MapBottomBar({super.key, this.currentPosition});

  @override
  State<MapBottomBar> createState() => _MapBottomBarState();
}

class _MapBottomBarState extends State<MapBottomBar> {
  int _lastSelectedIndex = 1; // 1 = maps page

  void _onItemTapped(int pressedIndex) {
    _handleNavigation(pressedIndex);

    setState(() {
      _lastSelectedIndex = pressedIndex;
    });
  }

  void _handleNavigation(int pressedIndex) {
    // Only map page is created so far.
    if (_isCreateNoteAction(pressedIndex)) {
      _navigateToCreateNote();
      return;
    }
  }

  bool _isCreateNoteAction(int index) {
    return _lastSelectedIndex == 1 && index == 1;
  }

  void _navigateToCreateNote() {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => BlocProvider(
          create: (_) => locator<NoteCubit>(),
          child: CreateNotePage(
            latitude: widget.currentPosition!.latitude,
            longitude: widget.currentPosition!.longitude,
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: <BottomNavigationBarItem>[
        BottomNavigationBarItem(icon: Icon(Icons.person), label: 'Profile'),
        BottomNavigationBarItem(
          icon: Icon(_lastSelectedIndex != 1 ? Icons.map : Icons.add),
          label: _lastSelectedIndex != 1 ? 'Map' : 'Create Note',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.notifications),
          label: 'Notifications',
        ),
      ],
      currentIndex: _lastSelectedIndex,
      selectedItemColor: Colors.amber,
      onTap: _onItemTapped,
    );
  }
}