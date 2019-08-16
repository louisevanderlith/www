import 'dart:async';
import 'dart:html';

class Carousel {
  DivElement _main;
  int _index;
  int _max;
  List<DivElement> _other;
  Duration _delay;

  Carousel(String mainElem, Duration delay) {
    _index = 0;
    _delay = delay;
    _main = querySelector(mainElem);

    if (_main != null) {
      _other = new List<DivElement>();

      var hasMore = true;
      var i = 0;

      do {
        final card = querySelector("#card${i}");
        hasMore = card != null;

        if (hasMore) {
          _other.add(card);
        }
        i++;
      } while (hasMore);

      _max = _other.length;
      new Timer.periodic(_delay, Move);
    }
  }

  void Move(Timer t) {
    _index++;
    if (_index == _max) {
      _index = 0;
    }

    for (var i = 0; i < _max; i++) {
      if (i == _index) {
        _other[i].classes.remove("hidden");
      } else {
        _other[i].classes.add("hidden");
      }
    }
  }
}
