$("#agree").bind("click", function(e){
    first = $("#v01").is(":checked");
    second = $("#v02").is(":checked");
    if(first==true && second==true) {
        e.preventDefault();
        location.href='join_continue.html';
    } else {
        alert("이용약관과 개인정보 수집 및 이용에 대한 안내 모두 동의해주세요.");
    }
});

$("#disagree").bind("click", function(e){
    e.preventDefault();
    location.href = "index.html";
});