var obj;
$(function(){
	$.post("/admin/list/management",
		function(result){
			obj = result;
			console.log(result.length);
			if (result.length == 0) {
				$("#adminListTable").find('tbody')
									  .append($('<tr></tr>'))
									  	.append($('<td></td>'))
									  		.attr('colspan', '7')
											.css('font-size', '15pt')
											.text('현재 승인이 필요한 관리자가 없습니다.');
				return;
			}

			$("#adminListTable").find('thead').append($("<tr>")
											  .append($("<td>").text("가입일자"))
											  .append($("<td>").text("아이디"))
											  .append($("<td>").text("이름"))
											  .append($("<td>").text("이메일"))
											  .append($("<td>").text("휴대폰번호"))
											  .append($("<td>").text("생일"))
											  .append($("<td>"))
			);

			for(i=0;i<result.length;i++){
				var row = $("<tr>");
				row.append($("<td>").attr('rowspan', '2').text(result[i].JoinDate));
				row.append($("<td>").attr('rowspan', '2').text(result[i].Id));
				row.append($("<td>").attr('rowspan', '2').text(result[i].Name));
				row.append($("<td>").attr('rowspan', '2').text(result[i].Email));
				row.append($("<td>").attr('rowspan', '2').text(result[i].Phone));
				row.append($("<td>").attr('rowspan', '2').text(result[i].Birth));
				row.append($("<td>").append($("<input>").attr({type: "button", value: "승인", id: "admit", name: i, class: "search", onClick: 'adminResult('+i+', admit)'})));

				var row2 = $("<tr>");
				row2.append($("<td>").append($("<input>").attr({type: "button", value: "거절", id: "refuse", name: i, class: "search", onClick: 'adminResult('+i+', refuse)'})));
				$("#adminListTable").append(row).append(row2);
			}
		}
	);
});

function adminResult(index, doing){
	alert(obj[index]);
	console.log(obj[index]);
	console.log(doing);

	if (doing == "admit"){
		$.ajax({
			type: 'POST',
			url: '/admin/list/do',
			data: { typeDo: "admit",
					id: obj[index].Id },
			success: function(){
				alert("해당 관리자의 신청이 승인되었습니다.");
				window.location.reload();
			},
			error: function(){
				alert("일시적인 오류가 발생했습니다. 다시 시도해주세요.");
			}
		});
	} else if (doing == "refuse"){
		$.ajax({
			type: 'POST',
			url: '/admin/list/do',
			data: { typeDo: "refuse",
					id: obj[index].Id },
			success: function(){
				alert("해당 관리자의 신청이 거절되었습니다.");
				window.location.reload();
			},
			error: function(){
				alert("일시적인 오류가 발생했습니다. 다시 시도해주세요.");
			}
		});
	}
}