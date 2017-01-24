var activeEl = 0;
$(function() {
    var items = $('.btn-nav');
    $( items[activeEl] ).addClass('active');
    $( ".btn-nav" ).click(function() {
        $( items[activeEl] ).removeClass('active');
        $( this ).addClass('active');
        activeEl = $( ".btn-nav" ).index( this );
    });
});

$(function(){
    $.ajax({
        url: "/mypage/getinfo",
        dataType: "json",
        success: function(data){
            $("#UserId").text(data.UserId).css("float", "left");
            $("#UserName").text(data.Name).css("float", "left");
            $("input[name=Email]").val(data.Email);
            $("input[name=PhoneNumber]").val(data.Phone);
            $("input[name=Birth]").val(data.Birth);
        }
    });
});

$("#modifySubmit").click(function(){
    $.ajax({
        url: "/mypage/modify",
        data: { password: $("#Password").val(), 
                passwordAgain: $("#PasswordAgain").val(),
                email: $("#Email").val(),
                phone: $("#Phone").val(),
                birth: $("#Birth").val() },
        success: modifyResult
    });
});

$("#withDrawalSubmit").click(function(){
    $.ajax({
        type: "POST",
        url: "/mypage/withdrawal",
        data: { password: $("#checkPassword").val() },
        success: withDrawalResult
    });
});

$("#modifyCancel, #withDrawalCancel").click(function(){
    location.href="service.html";
});

function modifyResult(result){
    if (result == "idpassword") {
        alert("아이디와 비밀번호는 일치해서는 안됩니다.");
        $("#Password").focus();
    } else if (result == "password") {
        alert("비밀번호는 8~15자의 영소문자와 숫자의 조합입니다.");
        $("#Password").focus();
    } else if (result == "passwordAgain") {
        alert("비밀번호가 일치하지 않습니다.");
        $("#PasswordAgain").focus();
    } else if (result == "email") {
        alert("이메일 주소를 입력하지 않았거나 형식이 올바르지 않습니다.");
        $("#Email").focus();
    } else if (result == "phone") {
        alert("휴대폰 번호를 입력하지 않았거나 형식이 올바르지 않습니다.");
        $("#Phone").focus();
    } else if (result == "birth") {
        alert("입력하신 생년월일의 형식이 올바르지 않습니다.");
        $("#Birth").focus();
    } else if (result == "success") {
        alert("회원정보 수정이 완료되었습니다.");
        $("#Password").val("");
        $("#PasswordAgain").val("");
    }
    event.preventDefault();
}

function withDrawalResult(result){
    alert(result);
    if (result == "disMatched"){
        alert("비밀번호가 일치하지 않습니다.");
        event.preventDefault();
    } else if (result == "success"){
        alert("탈퇴가 완료되었습니다.");
        location.href="index.html";
    }
}