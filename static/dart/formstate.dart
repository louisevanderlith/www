import 'dart:html';

class FormState {
  FormElement _form;
  ButtonElement _sendBtn;

  FormState(String formID, String submitID) {
    _form = querySelector(formID);
    _sendBtn = querySelector(submitID);

    disableSubmit(true);
  }

  bool isFormValid() {
    return _form.checkValidity();
  }

  void disableSubmit(bool disable) {
    _sendBtn.disabled = disable;
  }
}
