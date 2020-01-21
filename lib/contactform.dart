import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/bodies/message.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/commsapi.dart';

class ContactForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextAreaElement _message;
  ParagraphElement _error;
  String _tomail;

  ContactForm(String idElem, String nameElem, String emailElem,
      String phoneElem, String messageElem, String submitBtn)
      : super(idElem, submitBtn) {
    FormElement _form = querySelector(idElem);
    _form.onBlur.listen(validateElement);
    
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _message = querySelector(messageElem);
    _error = querySelector("${idElem}Err");

    var submit = querySelector(submitBtn);
    _tomail = submit.dataset['to'];

    submit.onClick.listen(onSend);
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
      new Toast.success(title: "Success!", message: content['Data'],position: ToastPos.bottomLeft);
    } else {
       new Toast.error(title: "Error!", message: content['Error'],position: ToastPos.bottomLeft);
    }
  }
}
