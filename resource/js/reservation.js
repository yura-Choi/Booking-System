$(".onewayterm").css("display", "none");
$("input[name=bus-way]:radio").change(function(){
    busway = document.getElementsByName('bus-way');
    if(busway[0].checked==true){
      	$(".onewayterm").css("display", "none");
      	$("#roundterm").css("display", "");
    } else if(busway[1].checked==true){
      	$(".onewayterm").css("display", "");
      	$("#roundterm").css("display", "none");
    }
});
$("input[name=bus-start]:radio").change(function(){
    busStart = document.getElementsByName('bus-start');
    if(busStart[0].checked==true){
        $("#resortResult").css("display", "");
        $("#seoulResult").css("display", "none");
    } else if (busStart[1].checked==true){
        $("#resortResult").css("display", "none");
        $("#seoulResult").css("display", "");
    }
});

$("#datepicker2").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        busStart = document.getElementsByName('bus-start');
        if (busStart[0].checked==true){
            $("#start_resort_date").text(dateText);
        } else if (busStart[1].checked==true){
            $("#start_seoul_date").text(dateText);
        }
    }
});

$(".btn_red").click(function(){
    $("#start_resort_place").text($(this).attr("value"));
    $("#start_resort_time").text($(this).text());
});

$(".btn_blue").click(function(){
    $("#start_seoul_place").text($(this).attr("value"));
    $("#start_seoul_time").text($(this).text());
});

$("#continueReservation").click(function(){
    busway = document.getElementsByName('bus-way');
    if (busway[0].checked){
        if ($("#start_resort_place").text()==""){
            alert("리조트행 셔틀버스를 선택해주세요");
            return;
        } else if ($("#start_seoul_place").text()==""){
            alert("서울행 셔틀버스를 선택해주세요");
            return;
        } else if ( $("#start_resort_date").text()==""){
            alert("출발일을 선택해주세요");
            return;
        } else if ( $("#start_seoul_date").text()==""){
            alert("도착일을 선택해주세요");
            return;
        }
        $.post("/reservation/next", { busType: "round",
                                      startDate: $("#start_resort_date").text(),
                                      startPlace: $("#start_resort_place").text(),
                                      startTime: $("#start_resort_time").text(),
                                      endDate: $("#start_seoul_date").text(),
                                      endPlace: $("#start_seoul_place").text(),
                                      endTime: $("#start_seoul_time").text()},
            function(result){
                if (result == "success"){
                    location.href="reservation_last.html";
                } else if (result == "error"){
                    alert("진행하는데 일시적인 오류가 발생하였습니다. 다시 시도해주세요.");
                }
            }
        );
    } else if (busway[1].checked){
        busToWhere = document.getElementsByName('bus-start');
        if (busToWhere[0].checked){
            if ($("#start_resort_place").text()==""){
                alert("리조트행 셔틀버스를 선택해주세요");
                return;
            } else if ( $("#start_resort_date").text()==""){
                alert("출발일을 선택해주세요");
                return;
            }
            $.post("/reservation/next", { busType:"resort",
                                          startDate: $("#start_resort_date").text(),
                                          startPlace: $("#start_resort_place").text(),
                                          startTime: $("#start_resort_time").text() },
                function(result){
                    if (result == "success"){
                        location.href="reservation_last.html";
                    } else if (result == "error"){
                        alert("진행하는데 일시적인 오류가 발생하였습니다. 다시 시도해주세요.");
                    }
                }
            );
        } else if (busToWhere[1].checked){
            if ($("#start_seoul_place").text()==""){
                alert("서울행 셔틀버스를 선택해주세요");
                return;
            } else if ( $("#start_seoul_date").text()==""){
                alert("출발일을 선택해주세요");
                return;
            }
            $.post("/reservation/next", { busType: "seoul",
                                          endDate: $("#start_seoul_date").text(),
                                          endPlace: $("#start_seoul_place").text(),
                                          endTime: $("#start_seoul_time").text() },
                function(result){
                    if (result == "success"){
                        location.href="reservation_last.html";
                    } else if (result == "error"){
                        alert("진행하는데 일시적인 오류가 발생하였습니다. 다시 시도해주세요.");
                    }
                }
            );
        } else {
            alert("탑승하실 방향을 선택해주세요.");
        }
    } else {
        alert("셔틀버스 타입을 선택해주세요.");
    }
});