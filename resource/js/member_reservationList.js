$(function(){
	$.post("/member/list/reserve",
		function(result){
			if (result.length == 0) {
				$("#reserveListTable").find('tbody')
									  .append($('<tr></tr>'))
									  	.append($('<td></td>'))
									  		.attr('colspan', '7')
											.css('font-size', '15pt')
											.text('예약된 내역이 없습니다.');
				return;
			}

			for(i=0;i<result.length;i++){
				if (result[i].BusType=="round"){
					var row = $("<tr>");
					row.append($("<td>").attr('rowspan', '2').text("왕복"));
					row.append($("<td>").attr('rowspan', '2').text(result[i].Member+"명"));
					row.append($("<td>").text("리조트행"));
					row.append($("<td>").text(result[i].ResortDate.String));
					row.append($("<td>").text(result[i].ResortPlace.String));
					row.append($("<td>").text(result[i].ResortTime.String));
					row.append($("<td>").attr('rowspan', '2')).append($("<input>").attr({type: "button", value: "예약취소", id: "cancelReservation"+i, class: "search"}));

					var row2 = $("<tr>");
					row2.append($("<td>").text("서울행"));
					row2.append($("<td>").text(result[i].SeoulDate.String));
					row2.append($("<td>").text(result[i].SeoulPlace.String));
					row2.append($("<td>").text(result[i].SeoulTime.String));

					$("#reserveListTable").append(row).append(row2);
				} else if (result[i].BusType=="resort"){
					var row = $("<tr>");
					row.append($("<td>").text("편도 - 리조트행"));
					row.append($("<td>").text(result[i].Member+"명"));
					row.append($("<td>").text("리조트행"));
					row.append($("<td>").text(result[i].ResortDate.String));
					row.append($("<td>").text(result[i].ResortPlace.String));
					row.append($("<td>").text(result[i].ResortTime.String));
					row.append($("<td>").append($("<input>").attr({type: "button", value: "예약취소", id: "cancelReservation"+i, class: "search"})));

					$("#reserveListTable").append(row);
				} else if (result[i].BusType=="seoul"){
					var row = $("<tr>");
					row.append($("<td>").text("편도 - 서울행"));
					row.append($("<td>").text(result[i].Member+"명"));
					row.append($("<td>").text("서울행"));
					row.append($("<td>").text(result[i].SeoulDate.String));
					row.append($("<td>").text(result[i].SeoulPlace.String));
					row.append($("<td>").text(result[i].SeoulTime.String));
					row.append($("<td>").append($("<input>").attr({type: "button", value: "예약취소", id: "cancelReservation"+i, class: "search"})));

					$("#reserveListTable").append(row);
				}
			}
		}
	);
});