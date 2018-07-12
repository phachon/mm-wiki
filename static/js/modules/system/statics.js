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
};