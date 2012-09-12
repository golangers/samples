$(function() {
    var pageName = document.location.pathname
    if (pageName == "/") {
        pageName = "/index.html"
    }
    
    $(".nav-collapse > ul > li > a[href='" + pageName + "']").parent("li").addClass("active");
});