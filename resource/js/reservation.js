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

$("#datepicker1-1").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        $("#start_resort_date").text(dateText);
    },
    defaultDate: "+1w",
    changeMonth: true,
    numberOfMonths: 1,
}).on("change", function(){
    $("#datepicker1-2").datepicker("option", "minDate", getDate(this));
});

$(function(){
    $("#datepicker1-1, #datepicker1-2, #datepicker2").datepicker("option", "showAnim", "slideDown");
    var d = new Date();
    var strDate = d.getFullYear()+"-"+(d.getMonth()+1)+"-"+(d.getDate()+1);
    $("#datepicker1-1, #datepicker2").datepicker("option", "minDate", strDate);
});

$("#datepicker1-2").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        $("#start_seoul_date").text(dateText);
    }
}).on("change", function(){
    $("#datepicker1-1").datepicker("option", "maxDate", getDate(this));
});

$("#datepicker2").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        busStart = document.getElementsByName('bus-start');
        if (busStart[0].checked==true){
            $("#start_resort_date").text(dateText);
            $("#start_seoul_date").text("");
            $("#start_seoul_place").text("");
            $("#start_seoul_time").text("");
        } else if (busStart[1].checked==true){
            $("#start_seoul_date").text(dateText);
            $("#start_resort_date").text("");
            $("#start_resort_place").text("");
            $("#start_resort_time").text("");
        }
    }
})

$(".btn_red").click(function(){
    $("#start_resort_place").text($(this).attr("value"));
    $("#start_resort_time").text($(this).text());
});

$(".btn_blue").click(function(){
    $("#start_seoul_place").text($(this).attr("value"));
    $("#start_seoul_time").text($(this).text());
});

$("#searchBusRound").click(function(){
    $.post("/reservation/search/round", {resortDate: $("#datepicker1-1").dateText, seoulDate: $("#datepicker1-2").dateText},
        function(result, status){
            if (result == ""){

            }
    });
});

$("#searchBusOne").click(function(){
    if (busStart[0].checked==true){
        $.post("/reservation/search/one", {startDate: $("#datepicker2").dateText},
            function(result, status){
                $(".btn_red").css("disable", "false");
        });
    } else if (busStart[1].checked==true){
        $.post("/reservation/search/one", {startDate: $("#datepicker2").dateText},
            function(result, status){
                $(".btn_blue").css("");
        });
    }
});

$("#continueReservation").click(function(){
    if (busway[0].checked){
        $.ajax({
            type: "POST",
            url: "/reservation/next/round",
            data: { type: "round",
                    startDate: $("#start_resort_date").text(),
                    startPlace: $("#start_resort_place").text(),
                    startTime: $("#start_resort_time").text(),
                    endDate: $("#start_seoul_date").text(),
                    endPlace: $("#start_seoul_place").text(),
                    endTime: $("#start_seoul_time").text()},
            success: function(){
                location.href="reservation_last.html";
            }
        })
    } else if (busway[1].checked){
        busToWhere = document.getElementsByName('bus-start');
        if (busToWhere[0].checked){
            $.ajax({
                type: "POST",
                url: "/reservation/next/one/resort",
                data: { type:"resort",
                        startDate: $("#start_resort_date").text(),
                        startPlace: $("#start_resort_place").text(),
                        startTime: $("#start_resort_time").text()},
                success: function(){
                    location.href="reservation_last.html";
                }
            });
        } else if (busToWhere[1].checked){
            $.ajax({
                type: "POST",
                url: "/reservation/next/one/seoul",
                data: { type: "seoul",
                        endDate: $("#start_seoul_date").text(),
                        endPlace: $("#start_seoul_place").text(),
                        endTime: $("#start_seoul_time").text() },
                success: function(){
                    location.href="reservation_last.html";
                }
            });
        } else {
            alert("탑승하실 방향을 선택해주세요.");
        }
    } else {
        alert("셔틀버스 타입을 선택해주세요.");
    }
});