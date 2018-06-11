/**
 * 公共类
 * Copyright (c) 2016 phachon@163.com
 */

$(function () {
	var isInFancybox = self != top;
	if(isInFancybox) {
		Common.cleanBody();
	}
});

var Common = {

	/**
	 * 是否在弹出层
	 */
	inPopup: false,

	/**
	 * 分页
	 * @param element
	 * @param pagedata
     */
	paginator: function(element, pagedata) {
	
		if (window["context"] == undefined) {
			if (!window.location.origin) {
				window.location.origin = window.location.protocol + "//" + window.location.hostname + (window.location.port ? ':' + window.location.port: '');
			}
			window["context"] = location.origin+"/V6.0";
		}
		
		var url = window.location.pathname;
		var host = window.location.origin;
		
		$(element).pagination({
			pageSize: pagedata[0].pageSize,
			pageIndex : pagedata[0].pageIndex,
			total: pagedata[0].total,
			firstBtnText:'首页',
			lastBtnText:'末页',
			showFirstLastBtn:true,
			showInfo: true,
			showJump: true,
			jumpBtnText:'跳转',
			infoFormat: '{start} ~ {end}条，共{total}条',
			noInfoText: '0 条数据',
		});

		$(element).on("pageClicked", function (event, data) {
			getRequest(data.pageIndex);
		}).on("jumpClicked", function (event, data) {
			getRequest(data.pageIndex);
		});

		//get请求
		function getRequest(index) {
			var urlQuerys = pagedata[0].urlQuery;
			var trueUrl = host + url + "?";
			for (key in urlQuerys) {
				trueUrl = trueUrl + "&"+key+"="+urlQuerys[key]+"&";
			}

			trueUrl = trueUrl + "page_index=" + index +"&page_size=" + pagedata[0].pageSize;
			window.location.href = trueUrl;
		}
	},

	/**
	 * 提示
	 * @param text
	 * @param url
	 * @param data
     * @param inPopup
     */
	confirm: function(text, url, data, inPopup) {
		if(inPopup) {
			Common.inPopup = true;
		}
		swal({
			title: text,
			type: "warning",
			showCancelButton: true,
			confirmButtonClass: "btn-danger",
			confirmButtonColor: "#D8534F",
			confirmButtonText: "是",
			cancelButtonText: "否",
			closeOnConfirm: false
		},
		function() {
			Common.submit(url, data, 'reload');
		});
	},

	/**
	 * ajax 提交
	 * @param url
	 * @param data
	 * @param location
	 * @param redirect
     */
	submit : function(url, data, location, redirect) {

		$.ajax({
			type : 'post',
			url : url,
			data : {'arr':data},
			dataType: "json",
			success : function(response) {
				if(response.code == 0) {
					swal({
						title: "操作失败",
						text: response.message,
						type: "error"
					});
				} else {
					swal({
						title: "操作成功",
						text: response.message,
						type: "success",
						showConfirmButton: false,
						timer: 2000
					});
				}

				Common.redirect(response.redirect);
			},
			error : function(response) {
				swal({
					title: "操作失败",
					text: response.message,
					type: "error"
				});
			}
		});
	},

	/**
	 * 跳转
	 * @param redirect
	 * @param inPopup
	 */
	redirect: function (redirect, inPopup) {
		if(inPopup != undefined) {
			Common.inPopup = inPopup;
		}
		//如果设置了跳转
		if(redirect) {
			setTimeout(function() {
				if(Common.inPopup) {
					parent.location.href = redirect;
				} else {
					location.href = redirect;
				}
			}, 2000);
			//重新刷新
			setTimeout(function() {
				if(Common.inPopup) {
					parent.location.reload();
				} else {
					location.reload();
				}
			}, 2000);
		}
	},

	/**
	 * 成功弹出框
	 * @param  text
	 * @return json
	 */
	successAlert: function (text) {
		swal({
			title: '操作成功',
			text: text,
			type: "success",
			showConfirmButton: false,
			timer: 2000
		});
	},

	/**
	 * 失败弹出框
	 * @param  text
	 * @return json
	 */
	errorAlert: function (text) {
		swal({
			title: '操作失败',
			text: text,
			type: "error",
			showConfirmButton: true,
			timer: 2000
		});
	},

	/**
	 * 警告弹出框
	 * @param  text
	 * @return json
	 */
	warningAlert: function (text) {
		swal({
			title: '警告',
			text: text,
			type: "warning",
			showConfirmButton: true,
			timer: 2000
		});
	},

	/**
	 * 删除弹出框
	 */
	confirmAlert: function (text) {
		var isConfirm = true;
		swal({
			title: "警告",
			text: text,
			type: "warning",
			showCancelButton: true,
			confirmButtonClass: "btn-danger",
			confirmButtonColor: "#DD6B55",
			confirmButtonText: "YES",
			cancelButtonText: "NO",
			closeOnConfirm: false,
		},
		function() {
			if (isConfirm) {
				swal("Deleted!", "Your imaginary file has been deleted.", "success");
			} else {
				swal("Cancelled", "Your imaginary file is safe :)", "error");
			}
		});
	},

	/**
	 * 根据 0~99 数字返回汉字 零~九十九
	 * @param  text
	 * @return string
	 */
	exchangeChinese: function (text) {
		text = parseInt(text);
		var chineses = ['零', '一', '二', '三', '四', '五', '六', '七', '八', '九', '十'];
		var remainder = text % 10;
		var number = parseInt(text / 10);

		if(number == 0) {
			return chineses[text];
		} else {
			var unit = chineses[remainder];//个位
			var decade = chineses[10];//十位
			if(number >= 2) {
				decade = chineses[number] + decade;
			}
			if(remainder == 0) {
				unit = '';
			}
			return decade + unit;
		}
	},

	/**
	 * 弹出浮层时清除 body 的样式
	 */
	cleanBody: function () {
		$("body").css('padding-top', 0);
		$("body").css('padding-bottom', 0);
	}
};
