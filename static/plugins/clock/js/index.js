// SNAPSVG.JS Clock Layout 

	// //clock 1
	// var clock1   = Snap("#clock1");
	// var hours1   = clock1.rect(79, 35, 2, 55).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
	// var minutes1 = clock1.rect(79, 20, 2, 70).attr({fill: "#282828", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
	// var seconds1 = clock1.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
	// var middle1 =   clock1.circle(80, 80, 2).attr({fill: "#ff6400"});
    //
	// //clock 2
	// var clock2   = Snap("#clock2");
	// var hours2   = clock2.rect(79, 35, 3, 45).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
	// var minutes2 = clock2.rect(79, 20, 3, 60).attr({fill: "#535353", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
	// var seconds2 = clock2.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
	// var middle2 =   clock2.circle(80, 80, 2).attr({fill: "#ff6400"});
    //
	// //clock 3
	// var clock3   = Snap("#clock3");
	// var hours3   = clock3.rect(79, 35, 3, 46).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
	// var minutes3 = clock3.rect(79, 20, 3, 60).attr({fill: "#535353", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
	// var seconds3 = clock3.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
	// var middle3 =   clock3.circle(81, 80, 3).attr({fill: "#535353"});

	//clock 4
	var clock4   = Snap("#clock4");
	var hours4   = clock4.rect(79, 35, 3, 55).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
	var minutes4 = clock4.rect(79, 20, 3, 70).attr({fill: "#535353", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
	var seconds4 = clock4.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
	var middle4 =   clock4.circle(81, 80, 3).attr({fill: "#535353"});

	// //clock 5
	// var clock5   = Snap("#clock5");
	// var hours5   = clock5.rect(79, 35, 3, 55).attr({fill: "#282828", transform: "r" + 10 * 30 + "," + 80 + "," + 80});
	// var minutes5 = clock5.rect(79, 20, 3, 70).attr({fill: "#535353", transform: "r" + 10 * 6 + "," + 80 + "," + 80});
	// var seconds5 = clock5.rect(80, 10, 1, 80).attr({fill: "#ff6400"});
	// var middle5 =   clock5.circle(80, 80, 2).attr({fill: "#ff6400"});


	// CLOCK Timer
	var updateTime = function(_clock, _hours, _minutes, _seconds) {
		var currentTime, hour, minute, second;
		currentTime = new Date();
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
	var updateSeconds = function(_clock, _seconds){
		var currentTime, second;
		currentTime = new Date();
		second = currentTime.getSeconds();

		if(second == 0){
			//got to 360deg at 60s
			second = 60;
		}else if(second == 1 && _seconds){
			//reset rotation transform(going from 360 to 6 deg)
			_seconds.attr({transform: "r" + 0 + "," + 80 + "," + 80});
		}
		if(_seconds){
			_seconds.attr({transform: "r" + second * 6 + "," + 80 + "," + 80});
		}
	};

	//update the clocks
	setInterval(function(){
		// updateTime(clock1, hours1, minutes1, seconds1);
		// updateTime(clock2, hours2, minutes2, seconds2);
		// updateTime(clock3, hours3, minutes3, seconds3);
		updateTime(clock4, hours4, minutes4, seconds4);
     	// updateTime(clock5, hours5, minutes5, seconds5);
	}, 1000);