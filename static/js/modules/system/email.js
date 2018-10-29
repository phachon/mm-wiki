// email module js

var Email = {

    // test send
    testSend : function (element, url, inPopup) {
        layer.prompt({title: '请输入收件人邮箱地址，多个以 ; 隔开', formType: 2}, function(text, index) {
            layer.close(index);
            $(element).attr('action', url);
            $(element).find("input[name='emails']").val(text);
            return Form.ajaxSubmit(element, inPopup);
        });
    }
};