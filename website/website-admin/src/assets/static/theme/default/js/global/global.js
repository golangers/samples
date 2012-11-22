function getUrlVars() {
    var vars = [],
        hash;
    var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&');
    for (var i = 0; i < hashes.length; i++) {
        hash = hashes[i].split('=');
        vars.push(hash[0]);
        vars[hash[0]] = hash[1];
    }
    return vars;
}

$(function(){
    $(".nav-collapse > ul > li > a[href='"+document.location.pathname+"']").parent("li").addClass("active");

    $("#message").bind('close',function(){
        $(this).hide().attr("class","alert").find("span").html("");
        return false;
    });
});