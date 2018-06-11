/**
 * main
 * Copyright (c) 2018
 */

/**
 * 调整工作区尺寸
 */
// function resizeContentHeight() {
// 	var mainHeight = document.body.clientHeight - 115;
// 	$('#menuFrame').height(mainHeight);
// }

function initMainLine() {
	var mainContent = document.getElementById("main-content");
	var mainLeft = document.getElementById("main-left");
	var mainRight = document.getElementById("main-right");
	var mainLine = document.getElementById("main-line");
	mainLine.onmousedown = function (e) {
		var disX = (e || event).clientX;
		mainLine.left = mainLine.offsetLeft;
		document.onmousemove = function (e) {
			var iT = mainLine.left + ((e || event).clientX - disX);
			var e = e || window.event, tarnameb = e.target || e.srcElement;
			var maxT = mainContent.clientWidth - mainLine.offsetWidth;
			mainLine.style.margin = 0;
			iT < 0 && (iT = 0);
			iT > maxT && (iT = maxT);
			mainLine.style.left = mainLeft.style.width = iT + "px";
			mainRight.style.width = mainContent.clientWidth - iT + "px";
			return false
		};
		document.onmouseup = function () {
			document.onmousemove = null;
			document.onmouseup = null;
			mainLine.releaseCapture && mainLine.releaseCapture()
		};
		mainLine.setCapture && mainLine.setCapture();
		return false
	};
}

function updateMainLeft(left) {
	var mainContent = document.getElementById("main-content");
	var mainLeft = document.getElementById("main-left");
	var mainRight = document.getElementById("main-right");
	var mainLine = document.getElementById("main-line");

	mainLine.left = mainLine.offsetLeft;
	var iT = left;
	var e = e || window.event, tarnameb = e.target || e.srcElement;
	var maxT = mainContent.clientWidth - mainLine.offsetWidth;
	mainLine.style.margin = 0;
	iT < 0 && (iT = 0);
	iT > maxT && (iT = maxT);
	mainLine.style.left = mainLeft.style.width = iT + "px";
	mainRight.style.width = mainContent.clientWidth - iT + "px";
}

function initPopover() {
    // webui-popover
    $("[data-toggle='web-popover']").webuiPopover({animation: 'pop',autoHide:3000});
}

function bindFancyBox() {
    $('a[name="create_space"]').each(function () {
        // $(this).fancybox({
        //     minWidth: 500,
        //     minHeight: 370,
        //     padding: 12,
        //     width: '65%',
        //     height: '48%',
        //     autoSize: false,
        //     type: 'iframe',
        //     href: $(this).attr('data-link')
        // });
        $(this).bind('click', function () {
            var content = $(this).attr("data-link");
            var height = "500px";
            var width = "1000px";
            layer.open({
                type: 2,
                // skin: Layers.skin,
                title: "ooooo",
                shadeClose: true,
                shade : 0.6,
                maxmin: true,
                area: [width, height],
                content: content,
                padding:"10px"
            });
        })
    });

    $('a[name="create_user"]').each(function () {
        $(this).fancybox({
            minWidth: 500,
            minHeight: 370,
            padding: 12,
            width: '65%',
            height: '48%',
            autoSize: false,
            type: 'iframe',
            href: $(this).attr('data-link')
        });
    });

    $('a[name="create_page"]').each(function () {
        $(this).fancybox({
            minWidth: 500,
            minHeight: 370,
            padding: 12,
            width: '65%',
            height: '48%',
            autoSize: false,
            type: 'iframe',
            href: $(this).attr('data-link')
        });
    });
}

$(window).resize(function() {
	// resizeContentHeight();
});

$(window).load(function() {
	initMainLine();
    initPopover();
    bindFancyBox();
});

