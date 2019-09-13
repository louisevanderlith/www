import 'package:WWW.APP/carousel.dart';
import 'package:WWW.APP/contactform.dart';

void main() {
  print("Running Default.Entry");
  new ContactForm("#frmContact", "#txtName", "#txtEmail", "#txtContact",
      "#txtMessage", "#btnSend");

  new Carousel("#syncarousel", new Duration(seconds:3));
}
