import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/bodies/key.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/commentapi.dart';
import 'package:mango_ui/bodies/comment.dart';

class Comments extends FormState {
  Key _objKey;
  String _objType;
  TextInputElement _text;

  Comments(String idElem, Key objKey, String objType)
      : super(idElem, "#btnComment") {
    _objKey = objKey;
    _objType = objType;
    _text = querySelector("#txtText");

    //data-itemKey="$key" data-itemType="Child"
    querySelector("#btnComment").onClick.listen(onCommentClick);
  }

  String get text {
    return _text.value;
  }

  void onCommentClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final data = new Comment(_objKey, text, _objType);
      var req = await createComment(data);
      
      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp['Data']);
      }
    }
  }
}
