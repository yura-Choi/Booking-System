$("#joinSubmit").click(function(e){
    if (checkInputValue()==false){
        $.post("/join/submit", {id: $("#UserId").val(), 
                                name: $("#UserName").val(), 
                                password: $("#Password").val(), 
                                passwordAgain: $("#PasswordAgain").val(), 
                                email: $("#Email").val(), 
                                phone: $("#Phone").val(), 
                                birth: $("#Birth").val()}, 
                function(result, status){
                    if (status == "success") {
                        joinNewAccount();
                    } else {
                        errorResult(result);
                    }
        });
    } else {
        e.preventDefault();
    }
});

var isValidId;
$(function(){
    isValidId=false;
});

$("#checkId").click(function(){
    $.ajax({
        type: "POST",
        url: "/join/checkid",
        data: {inputId:$("#UserId").val()},
        success: function(result){
            if (result == "alreadyExists"){
                alert("이미 존재하는 아이디입니다.");
                isValidId = false;
            } else if (result == "empty"){
                alert("아이디를 입력해주세요");
                isValidId = false;
            } else if (result == "success"){
                alert("사용가능한 아이디입니다.");
                isValidId = true;
            }
        }
    });
});

function joinNewAccount(){
    if (isValidId==false){
        alert("아이디 중복 체크를 해주세요.");
        $("#UserId").focus();
        return;
    } else if (isValidId==true){
        alert("회원가입이 완료되었습니다. 로그인 후 이용해주세요");
        location.href="index.html";
    }
}

function errorResult(result){
    if (result == "id") {
        alert("아이디는 8~10자의 영소문자와 숫자의 조합입니다.");
        $("#UserId").focus();
    } else if (isValidId==false){
        alert("아이디 중복 체크를 해주세요.");
        $("#UserId").focus();
    } else if (result == "name") {
        alert("이름을 입력하지 않았거나 길이가 20이상입니다.");
        $("#UserName").focus();
    } else if (result == "idpassword") {
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
    } else if (result == "dbfail") {
        alert("일시적인 오류가 발생하였습니다. 다시 시도해주세요.");
    } 
}

function checkChar(checkstr) {
	var numbercount = 0;
	var alphacount = 0;
    for (i=0; i <= checkstr.length; i++ ) {
        if (!isNaN(parseFloat(checkstr.charAt(i))) && isFinite(checkstr.charAt(i))) {
			numbercount++;
			continue;
        } else if (checkstr[i] == checkstr.charAt(i).toUpperCase()) {
			return true;
		}
		alphacount++;
	}
    if (numbercount >= checkstr.length || alphacount >= checkstr.length) {
		return true;
	}
	return false;
}

function checkInputValue(){
    var id = $("#UserId").val();
    var name = $("#UserName").val();
    var password = $("#Password").val();
    var passwordAgain = $("#PasswordAgain").val();
    var email = $("#Email").val();
    var phone = $("#Phone").val();
    var birth = $("#Birth").val();
    var emailValidate = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
    var phoneValidate = /^(?:(010-\d{4})|(01[1|6|7|8|9]-\d{3,4}))-(\d{4})$/;
    var birthValidate = /^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$/;
    if (id == "" || id.length<8 || id.length>10 || checkChar(id)){
        alert("아이디는 8~10자의 영소문자와 숫자의 조합입니다.");
        $("#UserId").focus();
        return true;
    } else if (name == "" || name.length >= 20){
        alert("이름을 입력하지 않았거나 길이가 20이상입니다.");
        $("#UserName").focus();
        return true;
    } else if (id == password) {
        alert("아이디와 비밀번호는 일치해서는 안됩니다.");
        $("#Password").focus();
        return true;
    } else if (password == "" || password.length<8 || password.length>15 || checkChar(password)){
        alert("비밀번호는 8~15자의 영소문자와 숫자의 조합입니다.");
        $("#Password").focus();
        return true;
    } else if (passwordAgain == "" || password != passwordAgain) {
        alert("비밀번호가 일치하지 않습니다.");
        $("#PasswordAgain").focus();
        return true;
    } else if (email == "" || !emailValidate.test(email)) {
        alert("이메일 주소를 입력하지 않았거나 형식이 올바르지 않습니다.");
        $("#Email").focus();
        return true;
    } else if (phone == "" || !phoneValidate.test(phone)) {
        alert("휴대폰 번호를 입력하지 않았거나 형식이 올바르지 않습니다.");
        $("#Phone").focus();
        return true;
    } else if (birth != "" && !birthValidate.test(birth)) {
        alert("입력하신 생년월일의 형식이 올바르지 않습니다.");
        $("#Birth").focus();
        return true;
    } else if (isValidId==false){
        alert("아이디 중복체크를 해주세요");
        return true;
    }
    return false;
}