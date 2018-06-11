/**
 * 表单提交类
 * Copyright (c) 2017 phachon
 */
var Form = {

	inPopup : false,

	/**
	 * ajax 提交表单
	 * @param element
	 * @param inPopup
	 * @returns {boolean}
	 */
	ajaxSubmit: function (element, inPopup) {

		var submitButton = $("button[name='submit']");
		if(inPopup) {
			Form.inPopup = true;
		}

		function successNotify(message, data) {
			var title = '<strong>操作成功：</strong>';
			submitButton.notify(title + message, {
				position: "right",
				className: 'success',
                autoHideDelay: 2500
			})
		}

		function failedNotify(errorMessage, data) {
			var title = "<strong>操作失败：</strong>";
			submitButton.notify(title + errorMessage, {
				position: "right",
				className: 'error',
                autoHideDelay: 2500
			})
		}

		function response(result) {
			if(result.code == 0) {
				failedNotify(result.message, result.data);
			}
			if(result.code == 1) {
				successNotify(result.message, result.data);
			}

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