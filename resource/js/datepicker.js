$("#datepicker1-1").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        $("#start_resort_date").text(dateText);
    }
}).on("change", function(){
    $("#datepicker1-2").datepicker("option", "minDate", getDate(this));
});

$(function(){
    $("#datepicker1-1, #datepicker1-2, #datepicker2").datepicker("option", "showAnim", "slideDown");
    var d = new Date();
    var strDate = d.getFullYear()+"-"+(d.getMonth()+1)+"-"+(d.getDate()+1);
    $("#datepicker1-1, #datepicker1-2, #datepicker2").datepicker("option", "minDate", strDate);
});

$("#datepicker1-2").datepicker({
    dateFormat: 'yy-mm-dd',
    onSelect: function(dateText, inst){
        $("#start_seoul_date").text(dateText);
    }
}).on("change", function(){
    $("#datepicker1-1").datepicker("option", "maxDate", getDate(this));
});