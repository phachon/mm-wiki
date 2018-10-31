// email module js

var Email = {

    // test send
    testSend : function (element, url) {
        layer.prompt({title: '请输入收件人邮箱地址，多个以 ; 隔开', formType: 2}, function(text, index) {
            layer.close(index);
            // $(element).attr('action', url);
            $(element).find("input[name='emails']").val(text);
            Email.ajaxSubmit(element, url);
        });
    },

    ajaxSubmit: function(element, url) {

        /**
         * 成功信息条
         * @param message
         * @param data
         */
        function successBox(message, data) {
            Common.successBox(Form.failedBox, message)
        }

        /**
         * 失败信息条
         * @param message
         * @param data
         */
        function failed(message, data) {
            Common.errorBox(Form.failedBox, message)
        }

        /**
         * request success
         * @param result
         */
        function response(result) {
            //console.log(result)
            if (result.code == 0) {
                failed(result.message, result.data);
            }
            if (result.code == 1) {
                successBox(result.message, result.data);
            }
            $("body,html").animate({scrollTop:0},300);
        }

        var options = {
            dataType: 'json',
            url: url,
            success: response
        };

        $(element).ajaxSubmit(options);
        return false;
    }
};