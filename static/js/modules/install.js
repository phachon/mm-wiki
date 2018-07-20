/**
 * 安装页
 * Copyright (c) 2017 phachon@163.com
 */

var Install = {

	
	/**
	 * get install status
     * @param url
	 * @constructor
	 */
	GetInstallStatus: function (url) {
		$.ajax({
			type : 'post',
			url : url,
			data : {'arr':''},
			dataType: "json",
			success : function(response) {
				if(response.code == 0) {
					console.log(response.message);
					return false
				}
				if (response.data.status == "1") {
					//正在安装中
					$("#install_success").addClass("hidden");
					$("#install_failed").addClass("hidden");
					$("#install_load").removeClass("hidden");
				}
				if (response.data.status == "2") {
					// 完成
					if (response.data.is_success == "1") {
						//失败
						$("#install_load").addClass("hidden");
						$("#install_success").addClass("hidden");
						$("#install_failed [data-name='error_message']").text(response.data.result);
						$("#install_failed").removeClass("hidden");
					}
					if (response.data.is_success == "2") {
						//成功
						$("#install_load").addClass("hidden");
						$("#install_failed").addClass("hidden");
						var res = eval('(' + response.data.result + ')');
						$("#install_success span[data-name='success_env_cmd'] code").text("set CODEPUBENV="+res.env);
						$("#install_success span[data-name='success_run_cmd'] code").text(res.cmd);
						$("#install_success a[data-name='success_message']").text(res.url);
						$("#install_success a[data-name='success_message']").attr('href', res.url);
						$("#install_success").removeClass("hidden");
					}
				}

			},
			error : function(response) {
				console.log(response.message)
			}
		});
	},

    /**
     * install ajax submit
     * @param formElement
     * @param element
     * @returns {boolean}
     */
    Submit: function(formElement, element) {

        /**
         * failed
         * @param message
         */
        function failed(message) {
            layer.tips(message, element, {tips: 3});
        }

        /**
         * response
         * @param result
         */
        function response(result) {
            if (result.code == 0) {
                failed(result.message, result.data);
            }
            if (result.code == 1) {
                // successBox(result.message, result.data);
            }
            if (result.redirect.url) {
                var sleepTime = result.redirect.sleep || 3000;
                setTimeout(function() {
                    location.href = result.redirect.url;
                }, sleepTime);
            }
        }

        var options = {
            dataType: 'json',
            success: response
        };

        $(formElement).ajaxSubmit(options);

        return false;
    }
};