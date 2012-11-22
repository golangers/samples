$(function() {
    $("#create_parent_category").chosen({
        allow_single_deselect: true
    });

    $("#update_parent_category").chosen({
        allow_single_deselect: true
    });

    $("input[op=op_delete]").click(function() {
        var col = $(this).parent().siblings()
        var id = col.eq(0).text();
        var name = col.eq(1).text();
        var h3 = $("#deleteModal .modal-header h3");
        h3.html("是否确认删除 " + name + "(" + id + ") ?");
        $(this).parent().parent().attr("ready", "true");
        $("#deleteModal").modal("show");
    });

    $("#delete_sure").click(function() {
        var tr = $("tr[ready=true]");
        var id = tr.children().eq(0).text();
        $.ajax({
            url: "delete.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&id=" + id,
            success: function(json) {
                if (json.status == 0) {
                    tr.removeAttr("ready");
                    $("#message").show().addClass("alert-error").find("span").html(json.message);
                    $("#deleteModal").modal("hide");
                } else {
                    window.document.location.reload();
                }
            }
        });
    })

    $("input[op=op_update]").click(function() {
        var col = $(this).parent().siblings()
        var id = col.eq(0).text();
        var name = col.eq(1).text();
        var order = col.eq(4).text();
        var h3 = $("#updateModal .modal-header h3");
        h3.html("修改 " + name + "(编号: " + id + ")");
        $(this).parent().parent().attr("ready", "true");
        $("#update_category").val(name);
        $("#update_category_order").val(order);
        $("#update_parent_category").val(col.eq(2).text());
        $("#update_parent_category").trigger("liszt:updated");
        $("#updateModal").modal("show");
    });
    $('button[id=btn_cancel_update]').click(function(){
        $('#updateModal').modal('hide');
    });

    $("#btn_update").click(function() {
        var tr = $("tr[ready=true]");
        var id = tr.children().eq(0).text();
        var name = $("#update_category");
        var nameHelp = name.next(".help-inline");
        var nameParent = name.parents(".control-group");

        var nameVal = name.val();

        name.focus(function() {
            nameParent.removeClass("error");
            nameHelp.hide().html("");
        });

        if (nameVal == "") {
            nameParent.addClass("error");
            nameHelp.html("类别名不能为空").show();
        }

        var parentCate = $("#update_parent_category");
        var parentCateHelp = parentCate.next(".help-inline");
        var parentCateParent = parentCate.parents(".control-group");
        var parentCateVal = parentCate.val();

        var order = $("#update_category_order");
        var orderVal = order.val();

        $.ajax({
            url: "update.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&id=" + id + "&name=" + nameVal + "&parentCate="+parentCateVal + "&order=" + orderVal,
            success: function(json) {
                if (json.status == 0) {
                    tr.removeAttr("ready");
                    nameParent.addClass("error");
                    nameHelp.html(json.message).show();
                } else if (json.status == 1) {
                    window.document.location.reload();
                }
            }
        });
    })

    $('input[op=op_stop]').click(function() {
        var col = $(this).parent().siblings()
        var id = col.eq(0).text();
        var name = col.eq(1).text();
        var status = $(this).val();
        var h3 = $("#stopModal .modal-header h3");
        h3.html("是否确认" + status + " " + name + " ?");
        $("#stop_sure").val("确认" + status);
        $(this).parent().parent().attr("ready", "true");
        $("#stopModal").modal("show");

    });

    $("#stop_sure").click(function() {
        var tr = $("tr[ready=true]");
        var id = tr.children().eq(0).text();

        $.ajax({
            url: "stop.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&id=" + id,
            success: function(json) {
                if (json.status == 0) {
                    tr.removeAttr("ready");
                    $("#message").show().addClass("alert-error").find("span").html(json.message);
                    $("#stopModal").modal("hide");
                } else {
                    window.document.location.reload();
                }
            }
        });
    });

    $('input[id=create_module]').click(function() {
        $("#createModal").modal("show");
    });
    $('button[id=btn_cancel_create]').click(function(){
        $('#createModal').modal('hide');
    });



    $("#btn_create").bind("click", function() {
        var name = $("#create_category");
        var nameHelp = name.next(".help-inline");
        var nameParent = name.parents(".control-group");

        var nameVal = name.val();

        name.focus(function() {
            nameParent.removeClass("error");
            nameHelp.hide().html("");
        });

        if (nameVal == "") {
            nameParent.addClass("error");
            nameHelp.html("模块名不能为空").show();
        }

        var parentCate = $("#create_parent_category");
        var parentCateHelp = parentCate.next(".help-inline");
        var parentCateParent = parentCate.parents(".control-group");
        var parentCateVal = parentCate.val();

        var order = $("#category_order");
        var orderVal = order.val();

        $.ajax({
            url: "create.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + nameVal + "&parentCate="+parentCateVal + "&order=" + orderVal,
            success: function(json) {
                if (json.status == 0) {
                    nameParent.addClass("error");
                    nameHelp.html(json.message).show();
                } else if (json.status == 1) {
                    window.document.location.reload();
                }
            }
        });
    });
});