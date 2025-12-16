import 'package:flutter/material.dart';

class MapBottomBar extends StatefulWidget {
  const MapBottomBar({super.key});

  @override
  State<MapBottomBar> createState() => _MapBottomBarState();
}
class _MapBottomBarState extends State<MapBottomBar> {
  int _selectedIndex = 1; // maps page

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: <BottomNavigationBarItem>[
        BottomNavigationBarItem(icon: Icon(Icons.history), label: 'History'),
        BottomNavigationBarItem(
          icon: Icon(_selectedIndex != 1 ? Icons.map : Icons.add),
          label: _selectedIndex != 1 ? 'Map' : 'Add Note',
        ),
        BottomNavigationBarItem(icon: Icon(Icons.person), label: 'Profile'),
      ],
      currentIndex: _selectedIndex,
      selectedItemColor: Colors.amber,
      onTap: _onItemTapped,
    );
  }
}