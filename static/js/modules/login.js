var Login = {

	errorMessage: "#errorMessage",

	/**
	 * 登录
	 * @param element
	 */
	ajaxSubmit: function (element) {

		var usernameEle = $(element).find("input[name='username']");
		var passwordEle = $(element).find("input[name='password']");
		var submitEle = $(element).find("input[name='submit']");
		var name = usernameEle.val();
		var password = passwordEle.val();
		if(!name) {
			layer.tips(usernameEle.attr("placeholder"), usernameEle);
			return false;
		}
		if (!password) {
            layer.tips(passwordEle.attr("placeholder"), passwordEle);
			return false;
		}

		function success(messages, data) {
			var text = "";
			var failedText = messages;
			$(Login.errorMessage).removeClass('alert-danger');
			$(Login.errorMessage).addClass('alert-success');
			$(Login.errorMessage).removeClass('hidden');
			$(Login.errorMessage + ' strong ').html(text + failedText);
		}

		function failed(messages, data) {
			var text = "登陆失败：";
			var failedText = messages;
			$(Login.errorMessage).removeClass('alert-success');
			$(Login.errorMessage).addClass('alert-danger');
			$(Login.errorMessage).removeClass('hidden');
			$(Login.errorMessage + ' strong ').html(text + failedText);
		}

		function response(result) {
			// $(Login.errorMessage).addClass('hidden');
			if(result.code == 0) {
				layer.tips(result.message, submitEle);
				// failed(result.message, result.data);
			}
			if(result.code == 1) {
				// success(result.message, result.data);
                var content = '<i class="fa fa-smile-o"></i> '+result.message;
                layer.msg(content);
                if (result.redirect.url) {
                    var sleepTime = result.redirect.sleep || 3000;
                    setTimeout(function() {
                    	location.href = result.redirect.url;
                    }, sleepTime);
                }
			}
		}

		var options = {
			dataType: 'json',
			success: response
		};

		$(element).ajaxSubmit(options);
	}
};