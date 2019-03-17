import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'pathlookup.dart';

class ContactForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextAreaElement _message;

  ContactForm(String idElem, String nameElem, String emailElem,
      String phoneElem, String messageElem, String submitBtn)
      : super(idElem, submitBtn) {
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _message = querySelector(messageElem);

    querySelector(submitBtn).onClick.listen(onSend);
    registerValidation();
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

  void registerValidation() {
    _name.onBlur.listen((e) => {validate(e, _name)});
    _email.onBlur.listen((e) => {validate(e, _email)});
    _phone.onBlur.listen((e) => {validate(e, _phone)});
    _message.onBlur.listen((e) => {validateArea(e, _message)});
  }

  void validate(Event e, InputElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void validateArea(Event e, TextAreaElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend().then((obj) => {disableSubmit(false)});
    }
  }

  Future submitSend() async {
    var url = await buildPath("Comms.API", "message", new List<String>());
    var data = jsonEncode({
      "Body": message,
      "Email": email,
      "Name": name,
      "Phone": phone,
      "To": ""
    });

    var resp = await HttpRequest.requestCrossOrigin(url, method: "POST", sendData: data);

    print(resp);
  }
}
