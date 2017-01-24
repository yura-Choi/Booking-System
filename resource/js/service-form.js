$('.message a').click(function(){
    $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
});

$(function () {
    if (window.location == window.parent.location) {
        $('#fullscreen').html('<span class="glyphicon glyphicon-resize-small"></span>');
        $('#fullscreen').attr('href', 'http://bootsnipp.com/mouse0270/snippets/PbDb5');
        $('#fullscreen').attr('title', 'Back To Bootsnipp');
    }    

    $('#fullscreen').on('click', function(event) {
        event.preventDefault();
        window.parent.location =  $('#fullscreen').attr('href');
    });
    
    $('#fullscreen').tooltip();
    $('.navbar-toggler').on('click', function(event) {
        event.preventDefault();
        $(this).closest('.navbar-minimal').toggleClass('open');
    });

    $.ajax({
        url: "/service/getCurrentStatus", 
        success: function(result){
            console.log(result);
            if (result == "admin"){
                $("span[name=reservation]").text("관리자 확인");
            } else if (result == "member"){
                $("span[name=reservation]").text("셔틀버스 예약");
            }
        }
    });
});

$("#mypage").click(function(){
    location.href="mypage.html";  //마이페이지 (공용)
});

$("#reservation").click(function(){
    $.post("/service/getCurrentStatus", function(result){
        if (result == "admin"){
            $("#reservation").attr("href","admin_management.html");  //관리자 확인(관리자용)
        } else if (result == "member"){
            $("#reservation").attr("href","reservation.html");       //셔틀버스 예약(회원용)
        }
    })
});

$("#reservationList").click(function(){
    $.post("/service/getCurrentStatus", function(result){
        console.log(result);
        if (result == "admin"){
            $("#List").attr("href", "admin_reservationList.html");    //예약된 목록 확인(관리자용)
        } else if (result == "member"){
            $("#List").attr("href", "member_reservationList.html");   //예약된 목록 확인(회원용)
        }
    })
});