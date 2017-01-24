$().ready(function () {
    $('#formid').on('keyup keypress', function (e) {
        var keyCode = e.keyCode || e.which;
        if (keyCode === 13) {
            e.preventDefault();
            return false;
        }
    });
});

$("#submitAdminCode").click(function () {
    $.ajax({
        type: "POST",
        dataType: "html",
        url: "/admincheck/value",
        data: { adminCheck: $("#checkAdminCode").val() },
        success: goNextJoinStep,
        error: errorResult
    });
});

function goNextJoinStep() {
    alert("관리자인증코드가 일치합니다. 관리자 회원가입을 진행합니다.");
    window.location.href = "join_inputinfo.html";
    return false;
}
function errorResult() {
    alert("관리자인증코드가 올바르지 않습니다.");
    return false;
}