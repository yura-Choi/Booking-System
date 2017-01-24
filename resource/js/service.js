$(function(){
    $.post("/service/getCurrentStatus",
        function(result, status){
            $.post("/index/getCookie",
                function(resultSession, statusSession){
                    if (resultSession == "") {
                        $("#navlogin").text("Login")
                                      .attr("href", "login.html");
                    } else if (result=="admin") {
                        $("#admin").css("display", "inline");
                        $("#member").css("display", "none");
                        $("#navlogin").text("Logout")
                                      .attr("href", "");
                    } else if (result=="member"){
                        $("#admin").css("display", "none");
                        $("#member").css("display", "inline");
                        $("#navlogin").text("Logout")
                                      .attr("href", "");
                    }
                }
            );
        }
    );
});

$("#mypage").click(function(){
    $.post("/index/getCookie", 
        function(result){
            if(result != ""){
                location.href="mypage.html";
            } else if (result == ""){
                alert("로그인 후 이용가능합니다.");
                location.href="login.html";
            }
        }
    );
});

$("#reservationBus").click(function(){
    $.post("/index/getCookie", 
        function(result){
            if(result != ""){
                location.href="reservation.html";
            } else if (result == ""){
                alert("로그인 후 이용가능합니다.");
                location.href="login.html";
            }
        }
    );
});

$("#showReservationInfo").click(function(){
    $.post("/index/getCookie", 
        function(result){
            console.log(result);
            if(result != ""){
                location.href="member_reservationList.html";
            } else if (result == ""){
                alert("로그인 후 이용가능합니다.");
                location.href="login.html";
            }
        }
    );
});

$("#managementCheck").click(function(){
    location.href="admin_reservationList.html";
});

$("#showBusInfo").click(function(){
    location.href="member_reservationList.html";
});