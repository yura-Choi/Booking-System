$("#generalbtn").bind("click", function(e){
    $.post("/join/select/type",
        function(){
            location.href="join_inputinfo.html";
    });

});

$("#administratorbtn").bind("click", function(e){
    e.preventDefault();
    window.location="join_administrator.html";
});