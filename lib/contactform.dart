import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_comms/bodies/message.dart';
import 'package:mango_comms/commsapi.dart';
import 'package:mango_ui/formstate.dart';

class ContactForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextAreaElement _message;

  ContactForm(String idElem, String submitBtn) : super(idElem, submitBtn) {
    _name = querySelector("#txtContactName");
    _email = querySelector("#txtContactEmail");
    _phone = querySelector("#txtContactContactNo");
    _message = querySelector("#txtContactBody");

    var submit = querySelector(submitBtn);

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

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend();
    }
  }

  submitSend() async {
    var data = new Message(message, email, name, phone);
    var req = await sendMessage(data);

    if (req.status == 200) {
      var content = jsonDecode(req.response);
      new Toast.success(
          title: "Success!",
          message: content['Data'],
          position: ToastPos.bottomLeft);
    } else {
      new Toast.error(
          title: "Error!",
          message: req.response,
          position: ToastPos.bottomLeft);
    }
  }
}
