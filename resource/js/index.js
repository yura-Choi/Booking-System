$("#title").text("셔틀버스 예약 시스템")
		   .css("color", "white")
		   .css("font-size", "35pt")
		   .appendTo("body");
$("#loginbtn").text("Login")
			  .addClass("btn btn-primary btn-lg outline")
			  .css("color", "white")
              .attr("onclick", "location.href='login.html'")
			  .appendTo("body"); 
$("#joinbtn").text("Join")
			 .addClass("btn btn-primary btn-lg outline")
			 .css("color", "white")
             .attr("onclick", "location.href='join.html'")
			 .appendTo("body");
$("#explanation").html("<center>본 사이트는 회원제로 운영되고 있습니다.<br>로그인을 하셔야 이용가능합니다.</center>")
				 .css("color", "white")
				 .appendTo("body");

$("#loginbtn").click(function(){
	if ($("#loginbtn").text()=="Logout"){
		$.post("/index/logout",
			function(){
				location.reload();
		});
	} else if ($("#loginbtn").text()=="Login"){
		location.href="login.html";
	}
});
