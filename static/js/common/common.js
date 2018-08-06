$(function () {
    $('[data-toggle="tooltip"]').tooltip();
    // $("[data-toggle='web-popover']").webuiPopover({animation: 'pop',autoHide:3000});
});

var Common = {

	/**
	 * ajax submit
	 * @param url
	 * @param data
	 */
	ajaxSubmit : function (url, data) {
		var jsonData = {};
		if (data !== "" && data !== undefined)  {
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
					Layers.failedMsg(response.message, function () {});
				} else {
					Layers.successMsg(response.message, Common.redirect(response.redirect.url));
				}
			},
			error : function(response) {
				Layers.failedMsg("server error!")
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
		}else {
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
	},

    secondsFormat: function (s) {
		var day = Math.floor( s/ (24*3600) ); // Math.floor()向下取整
		var hour = Math.floor( (s - day*24*3600) / 3600);
		var minute = Math.floor( (s - day*24*3600 - hour*3600) /60 );
		var second = s - day * 24 * 3600 - hour * 3600 - minute * 60;

		function formatStr(t) {
			if (parseInt(t) <= 0) {
				return "00"
			}
			if (0 < parseInt(t) && parseInt(t) < 10) {
				return "0"+t.toString()
			}
			return t
        }

		var timeRes = {
			d: formatStr(day),
			h: formatStr(hour),
			m: formatStr(minute),
			s: formatStr(second)
		};

		return timeRes
	}
};
