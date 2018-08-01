/**
 * Form.js 表单提交类
 * 依赖 jquery.form.js
 */

var Form = {

    /**
     * 提示 div
     */
    failedBox: '#failedBox',

    /**
     * 是否在弹框中
     */
    inPopup: false,

    /**
     * ajax submit
     * @param element
     * @param inPopup
     * @returns {boolean}
     */
    ajaxSubmit: function(element, inPopup) {

        if (inPopup) {
            Form.inPopup = true;
        }

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
            if (result.redirect.url) {
                var sleepTime = result.redirect.sleep || 3000;
                setTimeout(function() {
                    if (Form.inPopup) {
                        parent.location.href = result.redirect.url;
                    } else {
                        location.href = result.redirect.url;
                    }
                }, sleepTime);
            }
        }

        var options = {
            dataType: 'json',
            success: response
        };

        $(element).ajaxSubmit(options);

        return false;
    }
};