$(function(){

    $("input[name=new_class_name]").focus(function(){
        $(this).parent().next().children("span[op=class_edit]").fadeIn();
    });
    $("input[name=new_class_name]").blur(function(){
        $(this).parent().next().children("span[op=class_edit]").fadeOut();
    });

    $("span[op=img_delete]").click(function(){
        var id = $(this).parent().siblings().first().val();
        //console.log(id);
        $.ajax({
            url: "delete.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&id=" + id,
            success:function(json) {
               window.location = "/admin/index.html";
            }
        });
    });

    $("span[op=img_recover]").click(function(){
        var id = $(this).parent().siblings().first().val();
        //console.log(id);
        $.ajax({
            url: "recover.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&id=" + id,
            success:function(json) {
               window.location = "/admin/index.html";
            }
        });
    });


    $("span[op=class_delete]").click(function(){
        var classId = $(this).parent().siblings().first().val();
        //console.log(classId);
        $.ajax({
            url: "declass.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&classId=" + classId,
            success:function(json) {
               window.location = "/admin/index.html";
            }
        });
    });

    $("span[op=class_edit]").click(function(){
        var classId = $(this).parent().siblings().first().val();
        var className = $(this).parent().prev().children().first().val();
        console.log(classId);
        console.log(className);

        $.ajax({
            url: "edclass.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&classId=" + classId + "&className=" + className,
            success:function(json) {
               window.location = "/admin/index.html";
            }
        });
    });

    $("span[op=class_add]").click(function(){
        var className = $("#class_name").val();
        console.log(className);
        $.ajax({
            url: "adclass.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&className=" + className,
            success:function(json) {
               window.location = "/admin/index.html";
            }
        });
    });

});