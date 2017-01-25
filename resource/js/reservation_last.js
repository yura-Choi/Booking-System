$(function(){
    $.ajax({
        type: "POST",
        url: "/reservation/last",
        success: function(result){
			console.log(result.BusType);
            if (result.BusType == "round"){
                $("#resultResortPlaceTime").css("display", "");
                $("#resultSeoulPlaceTime").css("display", "");
                $("#busType").text("왕복");
                $("#start_resort_date").text(result.StartDate);
                $("#start_seoul_date").text(result.EndDate);
				$("#start_resort_place").text(result.StartPlace);
				$("#start_seoul_place").text(result.EndPlace);
				$("#start_resort_time").text(result.StartTime);
				$("#start_seoul_time").text(result.EndTime);
            } else if (result.BusType == "resort"){
                $("#resultResortPlaceTime").css("display", "")
                $("#resultSeoulPlaceTime").css("display", "none")
                $("#busType").text("편도(리조트행)");
                $("#start_resort_date").text(result.StartDate);
				$("#start_resort_place").text(result.StartPlace);
				$("#start_resort_time").text(result.StartTime);
            } else if (result.BusType == "seoul"){
                $("#resultResortPlaceTime").css("display", "none")
                $("#resultSeoulPlaceTime").css("display", "")
                $("#busType").text("편도(서울행)");
                $("#start_seoul_date").text(result.EndDate);
				$("#start_seoul_place").text(result.EndPlace);
				$("#start_seoul_time").text(result.EndTime);
            } else {
                alert("일시적인 오류가 발생했습니다. 다시 시도해주세요");
                location.href="reservation.html";
            }
        }
    });
});

$("#reservationSubmit").click(function(){
	$.ajax({
		type: "POST",
		url: "/reservation/submit",
		data: { member: $("#memberSelect option:selected").text()},
		success: function(){
			alert("예약이 완료되었습니다.");
			location.href="service.html";
		},
		error: function(){
            alert("일시적인 오류가 발생했습니다. 다시 시도해주세요");
		}
	});
});