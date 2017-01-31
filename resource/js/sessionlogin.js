$(function(){
    var cookie;
    $.post("/index/getCookie",
        function(result, status){
            if(result != ""){
                $("#navlogin, #loginbtn").text("Logout")
                                         .attr("href", "index.html");
            } else if(result == "") {
                $("#navlogin, #loginbtn").text("Login")
                                         .attr("href", "login.html");
            }
        }
    );
});

$("#navlogin, #loginbtn").click(function(){
    $.post("/index/logout",
        function(){
            location.href="index.html";
        }
    );                    
});