var obj;
$(function(){
	$("#form").css("max-width", "800px");
	$("#adminListTable").css("display", "none");
	$("#beforeShowList").css("display", "");

	$.post("/admin/list/management",
		function(result){
			obj = result;
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
											  .append($("<td>").text("생년월일"))
											  .append($("<td>"))
			);

			for(i=0;i<result.length;i++){
				var row = $("<tr>");
				row.append($("<td>").css("vertical-align", "middle").text(result[i].JoinDate));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Id));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Name));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Email));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Phone));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Birth));
				row.append($("<td>").css("vertical-align", "middle").append($("<input/>").css("height", "25px").css("margin-bottom", "5px").attr({type: "button", value: "승인", id: "admit", name: i, class: "search", onClick: 'adminResult('+i+', "admit")'})).append($("<br>")).append($("<input/>").css("height", "25px").attr({type: "button", value: "거절", id: "refuse", name: i, class: "search", onClick: 'adminResult('+i+', "refuse")'})));

				$("#adminListTable").append(row);
			}
		}
	);
});

function adminResult(index, doing){
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


$("#checkPassword").click(function(){
    $.post("/admin/list/checkpassword", {password: $("#inputPassword").val()},
        function(result){
            if (result == "correct"){
				$("#form").css("max-width", "1260px");
				$("#adminListTable").css("display", "");
				$("#beforeShowList").css("display", "none");
            } else if (result == "incorrect"){
                alert("비밀번호가 일치하지 않습니다.");
            }
    });
});