/**
 * 统计
 * Copyright (c) 2018 phachon@163.com
 */

var Statics = {

    GetSpaceDocsRank: function (element1, element2, url) {
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
                Morris.Bar({
                    element: element1,
                    data: response.data,
                    xkey: 'space_name',
                    ykeys: ['total'],
                    labels: ['文档数量'],
                    barRatio: 0.4,
                    xLabelAngle: 65,
                    hideHover: 'auto',
                    resize: true
                });
                var values = [];
                var count = 0;
                for (var i=0; i < response.data.length; i++) {
                    var value = {
                        value: response.data[i].total,
                        label: response.data[i].space_name
                    };
                    count += parseInt(response.data[i].total);
                    values.push(value)
                }
                Morris.Donut({
                    element: element2,
                    data: values,
                    formatter: function (x) { return Math.round((x/count)*100)+ "%"}
                }).on('click', function(i, row){
                    console.log(i, row);
                });
			},
			error : function(response) {
				console.log(response.message)
			}
		});
	},

    GetDocCountByTime: function (element, url) {
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
                if (response.data.length > 0) {
                    Morris.Line({
                        element: element,
                        data: response.data,
                        xkey: 'date',
                        ykeys: ['total'],
                        labels: ['新增文档数']
                    });
				}
            },
            error : function(response) {
                console.log(response.message)
            }
        });
    },

    GetCollectDocsRank: function (element, url) {
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
                if (response.data.length > 0) {
                    Morris.Bar({
                        element: element,
                        data: response.data,
                        xkey: 'document_name',
                        ykeys: ['total'],
                        labels: ['收藏数'],
                        barRatio: 0.4,
                        xLabelAngle: 65,
                        hideHover: 'auto',
                        resize: true
                    });
                }
            },
            error : function(response) {
                console.log(response.message)
            }
        });
    },

	GetServerStatus: function (url) {
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
				var cpu = response.data.cpu_used_percent;
				var memory = response.data.memory_used_percent;
				var disk = response.data.disk_used_percent;
				// cpu
				$(".cpu_text").each(function () {
					$(this).text(cpu+"%")
				});
				$("#cpu_progress").attr("aria-valuenow", cpu);
				$("#cpu_progress").attr('style', 'min-width: 2em; width: '+cpu+'%');

				// memory
				$(".memory_text").each(function () {
					$(this).text(memory+"%")
				});
				$("#memory_progress").attr("aria-valuenow", memory);
				$("#memory_progress").attr('style', 'min-width: 2em; width: '+memory+'%');

				// disk
				$(".disk_text").each(function () {
					$(this).text(disk+"%")
				});
				$("#disk_progress").attr("aria-valuenow", disk);
				$("#disk_progress").attr('style', 'min-width: 2em; width: '+disk+'%');

			},
			error : function(response) {
				console.log(response.message)
			}
		});
	},

    GetServerTime: function (element, url) {

        var clock4   = Snap("#"+element);
        var hours4   = clock4.rect(79, 35, 3, 55).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
        var minutes4 = clock4.rect(79, 20, 3, 70).attr({fill: "#535353", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
        var seconds4 = clock4.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
        var middle4 =   clock4.circle(81, 80, 3).attr({fill: "#535353"});

        var updateTime = function(serverTime, _clock, _hours, _minutes, _seconds) {

            var currentTime, hour, minute, second;
            currentTime = new Date(parseInt(serverTime) * 1000);
            second = currentTime.getSeconds();
            minute = currentTime.getMinutes();
            hour = currentTime.getHours();
            hour = (hour > 12)? hour - 12 : hour;
            hour = (hour == '00')? 12 : hour;

            if(second == 0){
                //got to 360deg at 60s
                second = 60;
            }else if(second == 1 && _seconds){
                //reset rotation transform(going from 360 to 6 deg)
                _seconds.attr({transform: "r" + 0 + "," + 80 + "," + 80});
            }
            if(minute == 0){
                minute = 60;
            }else if(minute == 1){
                _minutes.attr({transform: "r" + 0 + "," + 80 + "," + 80});
            }
            _hours.animate({transform: "r" + hour * 30 + "," + 80 + "," + 80}, 200, mina.elastic);
            _minutes.animate({transform: "r" + minute * 6 + "," + 80 + "," + 80}, 200, mina.elastic);
            if(_seconds){
                _seconds.animate({transform: "r" + second * 6 + "," + 80 + "," + 80}, 500, mina.elastic);
            }
        };

        var updateRuntime = function (runTime) {
            var timeRes = Common.secondsFormat(runTime);
            $("#run-days").text(timeRes.d);
            $("#run-hours").text(timeRes.h);
            $("#run-minutes").text(timeRes.m);
            $("#run-seconds").text(timeRes.s);
        };

        function getServerTime(url) {
            $.ajax({
                type : 'post',
                url : url,
                data : {},
                dataType: "json",
                success : function(response) {
                    // console.log(response.data.server_time);
                    // update lock time
                    updateTime(response.data.server_time, clock4, hours4, minutes4, seconds4);
                    // update run time
                    updateRuntime(response.data.run_time);
                },
                error : function(response) {
                    console.log("request error")
                }
            });
        }

        getServerTime(url);

        setInterval(function(){
            getServerTime(url);
        }, 1000);
    }
};