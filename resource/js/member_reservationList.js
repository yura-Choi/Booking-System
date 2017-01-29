var obj;
$(function(){
	$.post("/member/list/reserve",
		function(result){
			obj = result;
			if (result.length == 0) {
				$("#reserveListTable").find('tbody')
									  .append($('<tr></tr>'))
									  	.append($('<td></td>'))
									  		.attr('colspan', '7')
											.css('font-size', '13pt')
											.text('예약된 내역이 없습니다.');
				return;
			}

			$("#reserveListTable").find('thead').append($("<tr>")
											  .append($("<td>").text("구분"))
											  .append($("<td>").text("탑승인원"))
											  .append($("<td>").text("방향"))
											  .append($("<td>").text("탑승일자"))
											  .append($("<td>").text("탑승장소"))
											  .append($("<td>").text("탑승시간"))
											  .append($("<td>"))
			);

			for(i=0;i<result.length;i++){
				if (result[i].BusType=="왕복"){
					var row = $("<tr>");
					row.append($("<td>").attr('rowspan', '2').css("vertical-align", "middle").text("왕복"));
					row.append($("<td>").attr('rowspan', '2').css("vertical-align", "middle").text(result[i].Member+"명"));
					row.append($("<td>").text("리조트행"));
					row.append($("<td>").text(result[i].ResortDate.String));
					row.append($("<td>").text(result[i].ResortPlace.String));
					row.append($("<td>").text(result[i].ResortTime.String));
					row.append($("<td>").attr('rowspan', '2').css("vertical-align", "middle").append($("<input>").attr({type: "button", value: "예약취소", id: "cancelReservation", name: i, class: "search", onClick: 'deleteRow('+i+')'})));

					var row2 = $("<tr>");
					row2.append($("<td>").text("서울행"));
					row2.append($("<td>").text(result[i].SeoulDate.String));
					row2.append($("<td>").text(result[i].SeoulPlace.String));
					row2.append($("<td>").text(result[i].SeoulTime.String));

					$("#reserveListTable").append(row).append(row2);
				} else if (result[i].BusType=="편도(리조트행)"){
					var row = $("<tr>");
					row.append($("<td>").text("편도 - 리조트행"));
					row.append($("<td>").text(result[i].Member+"명"));
					row.append($("<td>").text("리조트행"));
					row.append($("<td>").text(result[i].ResortDate.String));
					row.append($("<td>").text(result[i].ResortPlace.String));
					row.append($("<td>").text(result[i].ResortTime.String));
					row.append($("<td>").append($("<input/>").attr({type: "button", value: "예약취소", id: "cancelReservation", name: i, class: "search", onClick: 'deleteRow('+i+')'})));

					$("#reserveListTable").append(row);
				} else if (result[i].BusType=="편도(서울행)"){
					var row = $("<tr>");
					row.append($("<td>").text("편도 - 서울행"));
					row.append($("<td>").text(result[i].Member+"명"));
					row.append($("<td>").text("서울행"));
					row.append($("<td>").text(result[i].SeoulDate.String));
					row.append($("<td>").text(result[i].SeoulPlace.String));
					row.append($("<td>").text(result[i].SeoulTime.String));
					row.append($("<td>").append($("<input/>").attr({type: "button", value: "예약취소", id: "cancelReservation", name: i, class: "search", onClick: 'deleteRow('+i+')'})));

					$("#reserveListTable").append(row);
				}
			}
		}
	);
});

function deleteRow(index){
	$.ajax({
		type: 'POST',
		url: '/member/list/delete',
		data: { member: obj[index].Member,
			    busType: obj[index].BusType,
			    resortDate: obj[index].ResortDate.String,
				seoulDate: obj[index].SeoulDate.String,
				resortPlace: obj[index].ResortPlace.String,
				seoulPlace: obj[index].SeoulPlace.String,
				resortTime: obj[index].ResortTime.String,
				seoulTime: obj[index].SeoulTime.String },
		success: function(){
			alert("예약이 취소되었습니다.");
			window.location.reload();
		},
		error: function(){
			alert("일시적인 오류가 발생했습니다. 다시 시도해주세요");
			window.location.reload();
		}
	});
}