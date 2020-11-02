import 'dart:html';

import 'package:WWW.APP/contactform.dart';

void main() {
  new ContactForm("#frmContact", "#btnContactSubmit");

  window.onScroll.listen(scrollFunction);
}

void scrollFunction(event) {
  if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
    document.getElementById("toTop").style.display = "block";
  } else {
    document.getElementById("toTop").style.display = "none";
  }
}
