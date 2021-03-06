import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_comment/bodies/comment.dart';
import 'package:mango_comment/commentapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class Comments extends FormState {
  Key _objKey;
  num _objType;
  TextInputElement _text;
  HiddenInputElement _userImg;

  Comments(String idElem, Key objKey, num objType)
      : super(idElem, "#btnComment") {
    _objKey = objKey;
    _objType = objType;
    _text = querySelector("#txtText");
    _userImg = querySelector("#hdnUserImg");

    //data-itemKey="$key" data-itemType="Child"
    querySelector("#btnComment").onClick.listen(onCommentClick);
  }

  String get text {
    return _text.value;
  }

  String get userImage {
    return _userImg.value;
  }

  void onCommentClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final data = new Comment(_objKey, text, _objType, userImage);
      var req = await createComment(data);

      if (req.status == 200) {
        final content = jsonDecode(req.response);
        new Toast.success(
            title: "Success!", message: content, position: ToastPos.bottomLeft);
      } else {
        new Toast.error(
            title: "Error!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}
