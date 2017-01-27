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

$(".mypage").click(function(){
    $.post("/index/getCookie",
        function(result){
            alert(result);
            if(result == "admin" || result == "member"){
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
    $.post("/service/getCurrentAdminAdmit", 
        function(result){
            if (result != "A"){
                alert("관리자의 승인을 받아야 이용할 수 있습니다.");
            } else if (result == "A"){
                location.href="admin_management.html";
            }
    });
});

$("#showBusInfo").click(function(){
    $.post("/service/getCurrentAdminAdmit", 
        function(result){
            if (result != "A"){
                alert("관리자의 승인을 받아야 이용할 수 있습니다.");
            } else if (result == "A"){
                location.href="admin_memberList.html";
            }
    });
});