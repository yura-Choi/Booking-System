$("input[name=bus-way]:radio").change(function(){
	busway = document.getElementsByName('bus-way');
	if (busway[0].checked==true){
		$("#default").css("display", "none");
		$("#resortSelect").css("display", "");
		$("#seoulSelect").css("display", "none");
		$("#seoulSelect option:eq(0)").attr("selected", "selected");
	} else if (busway[1].checked==true){
		$("#default").css("display", "none");
		$("#resortSelect").css("display", "none");
		$("#resortSelect option:eq(0)").attr("selected", "selected");
		$("#seoulSelect").css("display", "");
	}
});

$(function(){
	$("#datepicker").datepicker("option", "showAnim", "slideDown");

	$("#memberListTable").find('thead').append($("<tr>")
										.append($("<td>").css("width", "200px").css("text-align", "center").text("탑승시간"))
										.append($("<td>").css("width", "200px").css("text-align", "center").text("탑승인원"))
										.append($("<td>").css("width", "200px").css("text-align", "center").text("이름"))
										.append($("<td>").css("width", "200px").css("text-align", "center").text("휴대폰번호"))
										.append($("<td>").css("width", "200px").css("text-align", "center").text("생년월일"))
	);
});
var date;
$("#datepicker").datepicker({
    dateFormat: 'yy-mm-dd',
	onSelect: function(dateText, inst){
		date = dateText;
	}
});

var obj;
$(".search").click(function(){
	if ($("input[name=bus-way]:checked").val()==undefined){
		alert("버스의 방향을 선택하세요.");
		return;
	} else if (date==null){
		alert("날짜를 선택하세요");
		return;
	} else if (!($("#resortSelect option:selected").text()!="버스를 선택해주세요" || $("#seoulSelect option:selected").text()!="버스를 선택해주세요")){
		alert("버스를 선택하세요");
		return;
	}
	$("#memberListTable > tbody").empty();
	$("#memberListTable #title").empty();

	$.post("/admin/list/reservemember",
		{ busway: $("input[name=bus-way]:checked").val(),
		  date: date,
		  resortBus: $("#resortSelect option:selected").text(),
		  seoulBus: $("#seoulSelect option:selected").text()},
		function(result){
			obj = result;
			if (result.length == 0) {
				$("#memberListTable").find('tbody')
									  .append($('<tr></tr>')
									  	.append($('<td></td>')
									  		.attr('colspan', '5')
											.css('font-size', '13pt')
									  		.text('해당 날짜에 예약한 회원이 없습니다.')));
				return;
			}

			

			for(i=0;i<result.length;i++){
				var row = $("<tr>");
				row.append($("<td>").css("vertical-align", "middle").css("width", "150px").text(result[i].Time));
				row.append($("<td>").css("vertical-align", "middle").css("width", "150px").text(result[i].Member+"명"));
				row.append($("<td>").css("vertical-align", "middle").css("width", "150px").text(result[i].Name));
				row.append($("<td>").css("vertical-align", "middle").css("width", "200px").text(result[i].Phone));
				row.append($("<td>").css("vertical-align", "middle").css("width", "200px").text(result[i].Birth));

				$("#memberListTable").append(row);
			}
		}
	);
});

