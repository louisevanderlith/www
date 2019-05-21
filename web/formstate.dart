import 'dart:html';

class FormState {
  FormElement _form;
  ButtonElement _sendBtn;

  FormState(String formID, String submitID) {
    _form = querySelector(formID);
    _sendBtn = querySelector(submitID);

    disableSubmit(true);

    _form.onKeyUp.listen(pressEnter);
  }

  bool isFormValid() {
    return _form.checkValidity();
  }

  void disableSubmit(bool disable) {
    _sendBtn.disabled = disable;
  }

  void registerFormElements(List<Element> elements) {
    for (var elem in elements) {
      elem.onBlur.listen((e) {
        validateElement(e, elem);
      });
    }
  }

  void validateElement(Event e, InputElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    disableSubmit(!isFormValid());
  }

  void pressEnter(KeyboardEvent e) {
    if (e.key != 'Enter') {
      return;
    }

    e.preventDefault();
    _sendBtn.click();
  }
}
