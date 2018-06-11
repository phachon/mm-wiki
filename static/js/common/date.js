(function($) {
	$.extend({
		myTime: {
			/**
			 * 当前时间戳
			 * @return <int>        unix时间戳(秒)
			 */
			CurTime: function(){
				return Date.parse(new Date())/1000;
			},
			/**
			 * 日期 转换为 Unix时间戳
			 * @param <string> 2014-01-01 20:20:20  日期格式
			 * @return <int>        unix时间戳(秒)
			 */
			DateToUnix: function(string) {
				var f = string.split(' ', 2);
				var d = (f[0] ? f[0] : '').split('-', 3);
				var t = (f[1] ? f[1] : '').split(':', 3);
				return (new Date(
						parseInt(d[0], 10) || null,
						(parseInt(d[1], 10) || 1) - 1,
						parseInt(d[2], 10) || null,
						parseInt(t[0], 10) || null,
						parseInt(t[1], 10) || null,
						parseInt(t[2], 10) || null
					)).getTime() / 1000;
			},
			/**
			 * 时间戳转换日期
			 * @param <int> unixTime    待时间戳(秒)
			 * @param <bool> isFull    返回完整时间(Y-m-d 或者 Y-m-d H:i:s)
			 * @param <int>  timeZone   时区
			 */
			UnixToDate: function(unixTime, isFull, timeZone) {
                function add0(m){return m<10?'0'+m:m }

                if (typeof (timeZone) == 'number')
				{
					unixTime = parseInt(unixTime) + parseInt(timeZone) * 60 * 60;
				}
				var time = new Date(unixTime * 1000);
				var ymdhis = "";
				ymdhis += add0(time.getUTCFullYear()) + "-";
				ymdhis += add0((time.getUTCMonth()+1)) + "-";
				ymdhis += add0(time.getUTCDate());
				if (isFull === true)
				{
					ymdhis += " " + add0(time.getUTCHours()) + ":";
					ymdhis += add0(time.getUTCMinutes()) + ":";
					ymdhis += add0(time.getUTCSeconds());
				}
				return ymdhis;
			}
		}
	});
})(jQuery); 