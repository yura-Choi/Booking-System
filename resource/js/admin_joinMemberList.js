$(function(){
	$.post("/admin/list/joinedmember",
		function(result){
			if (result.length == 0){
				$("#joinedMemberTable").find('tbody')
									  .append($('<tr></tr>'))
									  	.append($('<td></td>'))
									  		.attr('colspan', '7')
											.css('font-size', '15pt')
											.text('현재 가입된 회원이 없습니다.');
				return;
			}

			$("#joinedMemberTable").find('thead').append($("<tr>")
											  .append($("<td>").text("가입일자"))
											  .append($("<td>").text("아이디"))
											  .append($("<td>").text("이름"))
											  .append($("<td>").text("이메일"))
											  .append($("<td>").text("휴대폰번호"))
											  .append($("<td>").text("생년월일"))
			);

			for(i=0;i<result.length;i++){
				var row = $("<tr>");
				row.append($("<td>").css("vertical-align", "middle").text(result[i].JoinDate));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Id));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Name));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Email));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Phone));
				row.append($("<td>").css("vertical-align", "middle").text(result[i].Birth));

				$("#joinedMemberTable").append(row);
			}
		}
	);
});