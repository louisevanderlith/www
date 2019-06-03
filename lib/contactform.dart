import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'pathlookup.dart';

class ContactForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextAreaElement _message;
  ParagraphElement _error;

  ContactForm(String idElem, String nameElem, String emailElem,
      String phoneElem, String messageElem, String submitBtn)
      : super(idElem, submitBtn) {
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _message = querySelector(messageElem);
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

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend();
    }
  }

  submitSend() async {
    var url = await buildPath("Comms.API", "message", new List<String>());
    var data = jsonEncode({
      "Body": message,
      "Email": email,
      "Name": name,
      "Phone": phone,
      "To": ""
    });

    var resp = await HttpRequest.requestCrossOrigin(url,
        method: "POST", sendData: data);
    var content = jsonDecode(resp);

    if (content['Error'] != "") {
      _error.text = content['Error'];
    } else {
      window.alert(content['Data']);
    }
  }
}
