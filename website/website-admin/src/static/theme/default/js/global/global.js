$(function(){
    $(".nav-collapse > ul > li > a[href='"+document.location.pathname+"']").parent("li").addClass("active");

    $("#message").bind('close',function(){
        $(this).hide().attr("class","alert").find("span").html("");
        return false;
    });
});