$("#generalbtn").bind("click", function(){
    $.post("/join/select/type",
        function(result){
            if (result == "success"){
                event.preventDefault();
                window.location="join_inputinfo.html";
            }
    });

});

$("#administratorbtn").bind("click", function(){
    event.preventDefault();
    window.location="join_administrator.html";
});