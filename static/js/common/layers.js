var Layers = {

	/**
	 * 皮肤
	 */
	skin : 'default',

	/**
	 * success 提示信息框
	 * @param title
	 */
	success : function (title) {
		layer.alert(title+"<br/>", {
			title: "操作成功",
			icon: 1,
			skin: Layers.skin,
			closeBtn: 0
		})
	},

	/**
	 * error 提示信息框
	 * @param title
	 */
	error : function (title) {
		layer.alert(title, {
			title: "操作失败",
			icon: 2,
			skin: Layers.skin,
			closeBtn: 0
		})
	},
	
	failedMsg: function (info) {
		var content = '<h4><i class="glyphicon glyphicon-remove"></i> 操作失败 </h4>';
		content += info;
		layer.msg(content, function(){});
	},

	successMsg: function (info) {
		var content = '<h4><i class="glyphicon glyphicon-ok"></i> 操作成功 </h4>';
		content += info;
		layer.msg(content);
	},

	/**
	 * confirm 提示框
	 * @param title
	 * @param url
	 */
	confirm: function (title, url) {
		layer.confirm(title, {
			btn: ['是','否'],
			skin: Layers.skin
		}, function() {
			Common.ajaxSubmit(url)
		}, function() {

		});
	},

	/**
	 * bind iframe 窗
	 */
	bindIframe: function (element, title, height, width, url) {
		$(element).each(function () {
			height = height||"500px";
			width = width||"1000px";
			$(this).bind('click', function () {
				var content = url || $(this).attr("data-link");
				layer.open({
					type: 2,
					skin: Layers.skin,
					title: '<strong>'+title+'</strong>',
					shadeClose: true,
					shade : 0.6,
					maxmin: true,
					area: [width, height],
					content: content,
					padding:"10px"
				});
			})
		})
	}
};