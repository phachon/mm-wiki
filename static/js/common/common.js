var Common = {

	/**
	 * ajax submit
	 * @param url
	 * @param data
	 */
	ajaxSubmit : function (url, data) {
		var jsonData = {};
		if (data !== "" && data !== undefined)  {
			console.log(data);
			var values = data.split("&");
            for (var i = 0; i < values.length; i ++) {
                jsonData[values[i].split("=")[0]] = unescape(values[i].split("=")[1]);
            }
		}

		$.ajax({
			type : 'post',
			url : url,
			data : jsonData,
			dataType: "json",
			success : function(response) {
				if(response.code == 0) {
					Layers.failedMsg(response.message)
				} else {
					Layers.successMsg(response.message)
				}
				Common.redirect(response.redirect.url);
			},
			error : function(response) {
				Layers.failedMsg(response.message)
			}
		});
	},

	/**
	 * redirect
	 * @param redirect
	 */
	redirect: function (redirect) {
		if(redirect) {
			setTimeout(function() {
				location.href = redirect;
			}, 2000);
			setTimeout(function() {
				location.reload();
			}, 2000);
		}
	},

	/**
	 * 成功提示
	 * @param element
	 * @param message
	 */
	successBox: function (element, message) {
		$(element).html('');
		$(element).removeClass();
		$(element).addClass('alert alert-success');
		$(element).append('<a class="close" href="#" onclick="$(this).parent().hide();">×</a>');
		$(element).append('<strong><i class="glyphicon glyphicon-ok-circle"></i> 操作成功：</strong>');
		$(element).append(message);
		$(element).show();
	},

	/**
	 * 错误提示
	 * @param element
	 * @param message
	 */
	errorBox: function (element, message) {
		$(element).html('');
		$(element).removeClass('hide');
		$(element).addClass('alert alert-danger');
		$(element).append('<a class="close" href="#" onclick="$(this).parent().hide();">×</a>');
		$(element).append('<strong><i class="glyphicon glyphicon-remove-circle"></i> 操作失败：</strong>');
		$(element).append(message);
		$(element).show();
	},

	/**
	 * 警告提示
	 * @param element
	 * @param message
	 */
	warningBox: function (element, message) {
		$(element).html('');
		$(element).removeClass();
		$(element).addClass('alert alert-warning');
		$(element).append('<a class="close" href="#" onclick="$(this).parent().hide();">×</a>');
		$(element).append('<strong><i class="glyphicon glyphicon-volume-up"></i> 警告：</strong>');
		$(element).append(message);
		$(element).show();
	}
};
