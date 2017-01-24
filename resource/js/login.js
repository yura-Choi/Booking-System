$('.message a').click(function(){
    $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
});

$("#login").click(function(){
    if ($("input[name=isAdmin]:checked").val() == null) {
        alert("로그인 회원타입을 선택해주세요");
        return
    }
    $.post("/login", {id: $("#loginId").val(),
                      password: $("#loginPassword").val(),
                      isAdmin: $("input[name=isAdmin]:checked").val()},
            function(result, status, xhr){
                if (result == "WrongId"){
                    alert("아이디가 존재하지 않습니다.");
                } else if(result == "WrongPassword"){
                    alert("비밀번호가 일치하지 않습니다.");
                } else if(result == "loginRestriction"){
                    alert("비밀번호를 5회이상 잘못 입력하여 해당 계정의 로그인이 제한되었습니다. 이메일 인증 후 이용해주시기 바랍니다.");
                    checkAccount();
                } else if(result == "success") {
                    location.href="index.html";
                }
            });
});

$("#findId").click(function(){
    if ($("input[name=isAdminToFind]:checked").val() == null) {
        alert("찾으시려는 회원타입을 선택해주세요");
        return
    }
    $.ajax({
        type: "POST",
        dataType: "html",
        url: "/login/findid",
        data: { name: $("#idName").val(),
                email: $("#idEmail").val(),
                isAdmin: $("input[name=isAdminToFind]:checked").val()},
        success: findUserId,
        error: errorResult
    });
});

$("#findPassword").click(function(){
    if ($("input[name=isAdminToFind]:checked").val() == null) {
        alert("찾으시려는 회원타입을 선택해주세요");
        return
    }
    $.ajax({
        type: "POST",
        dataType: "html",
        url: "/login/findpassword",
        data: {name: $("#passwordName").val(),
               id: $("#passwordId").val(),
               email: $("#passwordEmail").val(),
               isAdmin: $("input[name=isAdminToFind]:checked").val()},
        success: findUserPassword,
        error: errorResult
    });
});

function findUserId(result){
    //회원 아이디 화면에 보여주기
    var showIdHint;
    var length=result.length;
    showIdHint=result.slice(0, 4);
    for(var i=4;i<length;i++){
        showIdHint+="*";
    }
    alert("아이디: " + showIdHint);
    window.location.reload(true);
}

function findUserPassword(){
    alert("해당 계정이 존재합니다.");
    //새롭게 설정할 비밀번호를 입력받고 설정하기
    location.href="setNewPassword.html";    
}
$("#passwordSetting").click(function(){
    $.post("/login/setpassword", {password: $("#setPassword").val(), passwordAgain: $("#setPasswordAgain").val()},
        function(result, status){
            if (result == "Mismatched"){
                alert("바꾸시려는 비밀번호를 한 번 더 입력해주세요");
                $("#setPasswordAgain").focus();
            } else if (result == "MatchedPrevPassword"){
                alert("현재 비밀번호와 일치합니다.");
                $("#setPassword").focus();
            } else if (result == "WrongPassword"){
                alert("비밀번호는 8~15자의 영소문자와 숫자의 조합입니다.");
                $("#setPassword").focus();
            } else if(result == "SameWithId"){
                alert("비밀번호는 아이디와 일치해서는 안됩니다.");
                $("#setPassword").focus();
            } else if (result == "ChangeSuccess"){
                alert("비밀번호 설정에 성공했습니다. 로그인 후 이용해주세요.");
                location.href="login.html";
            }
        }
    );
});

function errorResult(){
    alert("해당 회원 정보가 없습니다.");
    location.reload();
}

$("#sendCodeBtn").click(function(){
    if ($("#sendCodeBtn").attr("value") == "인증번호 발급") {
        $.post("/login/sendcode", function(result, status){
                if(result == "success"){
                    alert("인증번호가 발급되었습니다. 회원가입 시 입력하셨던 이메일주소로 인증번호를 확인해주세요.");
                } else if (result == "fail"){
                    alert("일시적인 오류로 인증번호 발급에 실패했습니다. 다시 시도해주세요");
                }
            }
        );
        $("#sendCodeBtn").attr("value","인증번호 확인");
    } else if ($("#sendCodeBtn").attr("value") == "인증번호 확인"){
        $.post("/login/checkcode", {code: $("#code").val()},
            function(result, status){
                if (result == "success"){
                    alert("인증에 성공했습니다. 로그인 후 이용해주세요.");
                    window.close();
                    location.reload();
                } else if (result == "WrongCertificationCode"){
                    alert("인증번호가 일치하지 않습니다.");
                }
            }
        );
    }
});

function checkAccount() { 
    window.open('../certificateAccount.html','','location=no, directories=no,resizable=no,status=no,toolbar=no,menubar=no, width=1000,height=400,left=0, top=0, scrollbars=no');
}