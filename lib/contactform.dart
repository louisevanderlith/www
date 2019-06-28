import 'dart:convert';
import 'dart:html';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/bodies/message.dart';
import 'package:mango_ui/services/commsapi.dart';

class ContactForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextAreaElement _message;
  ParagraphElement _error;
  String _tomail;

  ContactForm(String idElem, String nameElem, String emailElem,
      String phoneElem, String messageElem, String tomail, String submitBtn)
      : super(idElem, submitBtn) {
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _message = querySelector(messageElem);
    _tomail = tomail;
    _error = querySelector("${idElem}Err");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get name {
    return _name.value;
  }

  String get email {
    return _email.value;
  }

  String get phone {
    return _phone.value;
  }

  String get message {
    return _message.value;
  }

  String get to {
    return _tomail;
  }
  
  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend();
    }
  }

  submitSend() async {
    var data = new Message(message, email, name, phone, to);
    var req = await sendMessage(data);
    var content = jsonDecode(req.response);

    if (req.status == 200) {
      window.alert(content['Data']);
    } else {
      _error.text = content['Error'];
    }
  }
}
