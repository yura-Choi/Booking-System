$("#generalbtn").bind("click", function(){
    $.post("/join/select/type",
        function(result){
            if (result == "success"){
                location.href="join_inputinfo.html";
            }
    });

});

$("#administratorbtn").bind("click", function(){
    location.href="join_administrator.html";
});