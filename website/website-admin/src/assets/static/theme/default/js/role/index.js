$(function() {
    $("#create #scope").chosen({
        allow_single_deselect: true
    });
    $("#update #scope").chosen({
        allow_single_deselect: true
    });
    $("#add_role_user").chosen();

    $('#create_role').click(function() {
        $("#createModal").modal("show");
    });

    $("#btn_create").bind("click", function() {
        var name = $("input[name=rolename]");
        var nameHelp = name.next(".help-inline");
        var nameParent = name.parents(".control-group");

        var scope = $("#create #scope");
        var scopeHelp = scope.nextAll(".help-inline");
        var scopeParent = scope.parents(".control-group");

        var nameVal = name.val(),
            scopeVal = scope.val(),
            shouldReturn = false;

        name.focus(function() {
            nameParent.removeClass("error");
            nameHelp.hide().html("");
        });

        scope.focus(function() {
            scopeParent.removeClass("error");
            scopeHelp.hide().html("");
        });

        if (nameVal == "") {
            nameParent.addClass("error");
            nameHelp.html("角色名不能为空").show();
            shouldReturn = true;
        }

        if (scopeVal == "") {
            scopeParent.addClass("error");
            scopeHelp.html("作用域不能为空").show();
            shouldReturn = true;
        }

        if (shouldReturn) {
            return;
        }

        $.ajax({
            url: "create.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + nameVal + "&scope=" + scopeVal,
            success: function(json) {
                if (json.status == 0) {
                    nameParent.addClass("error");
                    nameHelp.html(json.message).show();
                } else {
                    window.document.location.reload();
                }
            }
        });
    });

    $('input[op=role_show_users]').click(function() {
        var col = $(this).parent().parent().children()
        var name = col.eq(0).text();
        var users = col.eq(1).attr("data").split(",");
        var h3 = $("#usersModal .modal-header h3");
        h3.html("设置拥有<" + name + ">角色的用户");
        $("#add_role_user").val(users);
        $("#add_role_user").trigger("liszt:updated");
        $(this).parent().parent().attr("ready", "true");
        $("#usersModal").modal("show");
    });

    $("#btn_add_users").bind("click", function() {
        var tr = $("tr[ready=true]");
        var name = tr.children().eq(0).text();
        var users = $("#add_role_user");
        var usersVal = users.val();
        if (usersVal == null) {
            usersVal = "";
        }

        $.ajax({
            url: "users.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + name + "&users=" + usersVal,
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

    current = 1;
    expand = true; //decide expand or not
    $("#action_models > li").hide();
    $("#models > li:nth-child(" + current.toString() + ")").addClass("current");
    $("#action_models > li:nth-child(" + current.toString() + ")").show();

    var choice = {};

    function init_choice() {
        choice = {};
        $("#models").children().each(function() {
            choice[$(this).text()] = {};
        });
        current = 1;
    }

    function clear_model() {
        $("#models > li > .icon-blank").each(function() {
            this_model = $(this).parent();
            model_num = this_model.parent().children().index(this_model) + 1;
            this_model.removeClass("chosen");
            abandon_actions($("#action_models > li:nth-child(" + model_num.toString() + ") > ul > li"));
        });
    }

    $('input[op=right]').click(function() {
        var col = $(this).parent().parent().children()
        var name = col.eq(0).text();
        var h3 = $("#rightModal .modal-header h3");
        h3.html("设置角色<" + name + ">的权限");
        var tr = $(this).parent().parent();
        tr.attr("ready", "true");
        $("#rightModal").modal("show");
        var json = $.parseJSON($(this).attr("modules"));
        init_choice();
        clear_model();
        if (json) {
            $.each(json,function(k,v){
                $("#models").children().each(function() {
                    var thisObj = $(this)
                    if (v["module"] == $(this).text()) {
                        thisObj.attr("choose","1");
                        $.each(v["actions"],function(ak,av){
                            model_num = thisObj.parent().children().index(thisObj) + 1;
                            this_actions = $("#action_models > li:nth-child(" + model_num.toString() + ") > ul > li");
                            $.each(this_actions, function(i,this_action){
                                if (av == $(this_action).text()) {
                                    choose_action($(this_action));
                                }
                            })
                        });
                    }
                });
            });
        }
    });

    $("#btn_set_right").bind("click", function() {
        var tr = $("tr[ready=true]");
        var name = tr.children().eq(0).text();
        var choose_models = {};
        $("#models > li[choose=1]").each(function () {
            var thisObj = $(this);
            if (choice[thisObj.text()]) {
                choose_models[thisObj.text()] = []
                $.each(choice[thisObj.text()],function(k,v){
                    if (v == true) {
                        choose_models[thisObj.text()].push(k);
                    }
                });
            }
        });

        $.ajax({
            url: "right.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + name + "&right=" + JSON.stringify(choose_models),
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

    function choose_action(this_action) {
        action_num = this_action.parent().children().index(this_action) + 1;
        model_num = this_action.parents("#action_models").children().index(this_action.parent().parent()) + 1;

        this_model = $("#models > li:nth-child(" + model_num.toString() + ")");
        choice[this_model.text()][this_action.text()] = true;
        this_action.addClass("chosen");

        if (this_model.attr("choose") != "1") {
            this_model.attr("choose","1");
        }

        if (!this_model.hasClass("chosen")) {
            this_model.addClass("chosen");
        }

        if (this_action.siblings().filter(".chosen").length == this_action.siblings().length) {
            this_model.addClass("chosen");
        }
    }

    function abandon_action(this_action) {
        action_num = this_action.parent().children().index(this_action) + 1;
        model_num = this_action.parents("#action_models").children().index(this_action.parent().parent()) + 1;

        this_model = $("#models > li:nth-child(" + model_num.toString() + ")");
        choice[this_model.text()][this_action.text()] = false;
        this_action.removeClass("chosen");

        if (this_action.siblings().filter(".chosen").length == 0) {
            this_model.removeClass("chosen");
            this_model.attr("choose","0");
        }
    }

    function choose_actions(this_actions) {
        num = this_actions.length;
        if (num >= 1) {
            this_actions.each(function() {
                choose_action($(this));
            });
        } else {
            choose_action(this_actions);
        }
    }

    function abandon_actions(this_actions) {
        num = this_actions.length;
        if (num >= 1) {
            this_actions.each(function() {
                abandon_action($(this));
            });
        } else {
            abandon_action(this_actions);
        }
    }

    function choose_model(this_model) {
        model_num = this_model.parent().children().index(this_model) + 1;
        this_model.addClass("chosen");
        this_actions = $("#action_models > li:nth-child(" + model_num.toString() + ") > ul > li");
        choose_actions(this_actions);
        this_model.attr("choose","1");
    }

    function abandon_model(this_model) {
        model_num = this_model.parent().children().index(this_model) + 1;

        this_model.removeClass("chosen");
        abandon_actions($("#action_models > li:nth-child(" + model_num.toString() + ") > ul > li"));
    }

    $("#models > li > .icon-blank").click(function() {
        expand = false;
        this_model = $(this).parent();
        model_num = $("ul#models > li").index(this_model) + 1;

        if (!this_model.hasClass("chosen")) {
            choose_model(this_model);
        } else {
            abandon_model(this_model);
        }
    });

    $(".actions > li > .icon-blank").click(function() {
        this_action = $(this).parent();
        if (!this_action.hasClass("chosen")) {
            choose_action(this_action);
        } else {
            abandon_action(this_action);
        }
    });

    $("#models > li").click(function() {
        if (expand == true) {
            $("#models > li:nth-child(" + current.toString() + ")").removeClass("current");
            $("#action_models > li:nth-child(" + current.toString() + ")").hide();
            current = $("ul#models > li").index($(this)) + 1;
            $("#models > li:nth-child(" + current.toString() + ")").addClass("current");
            $("#action_models > li:nth-child(" + current.toString() + ")").show();
        }
        expand = true;
    });

    $('input[op=role_stop]').click(function() {
        var col = $(this).parent().parent().children()
        var name = col.eq(0).text();
        var status = $(this).val();
        var h3 = $("#stopModal .modal-header h3");
        h3.html("是否确认" + status + " " + name + " ?");
        $("#stop_sure").val("确认" + status);
        $(this).parent().parent().attr("ready", "true");
        $("#stopModal").modal("show");
    });

    $("#stop_sure").click(function() {
        var tr = $("tr[ready=true]");
        var name = tr.children().eq(0).text();

        $.ajax({
            url: "stop.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + name,
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

    $("input[op=role_delete]").click(function() {
        var h3 = $("#deleteModal .modal-header h3");
        var col = $(this).parent().parent().children()
        var name = col.eq(0).text();
        h3.html("是否确认删除 " + name + " ?");
        $(this).parent().parent().attr("ready", "true");
        $("#deleteModal").modal("show");
    });

    $("#delete_sure").click(function() {
        var tr = $("tr[ready=true]");
        var name = tr.children().eq(0).text();

        $.ajax({
            url: "delete.html",
            type: "POST",
            dataType: "json",
            data: "ajax=&name=" + name,
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
    });

    $('input[op=role_update]').click(function() {
        var h3 = $("#updateModal .modal-header h3");
        var col = $(this).parent().parent().children()
        var name = col.eq(0).text();
        $("input[name=oldname]").val(name);
        $("#update #scope").val(col.eq(2).attr("data"));
        $("#update #scope").trigger("liszt:updated");
        h3.html("是否确认修改 " + name + " 的角色名?");
        $(this).parent().parent().attr("ready", "true");
        $("#updateModal").modal("show");

    });

    $("#btn_update").bind("click", function() {
        var name = $("input[name=name]");
        var oldname = $("input[name=oldname]").val();
        var nameHelp = name.next(".help-inline");
        var nameParent = name.parents(".control-group");

        var scope = $("#update #scope");
        var scopeHelp = scope.nextAll(".help-inline");
        var scopeParent = scope.parents(".control-group");

        var nameVal = name.val(),
            scopeVal = scope.val(),
            shouldReturn = false;

        name.focus(function() {
            nameParent.removeClass("error");
            nameHelp.hide().html("");
        });

        scope.focus(function() {
            scopeParent.removeClass("error");
            scopeHelp.hide().html("");
        });

        if (nameVal == oldname) {
            nameParent.addClass("error");
            nameHelp.html("请修改为新的角色名").show();
            shouldReturn = true;
        }

        if (scopeVal == "") {
            scopeParent.addClass("error");
            scopeHelp.html("作用域不能为空").show();
            shouldReturn = true;
        }

        if (shouldReturn) {
            return;
        } else {
            $.ajax({
                url: "update.html",
                type: "POST",
                dataType: "json",
                data: "ajax=&oldname=" + oldname + "&name=" + nameVal + "&scope=" + scopeVal,
                success: function(json) {
                    if (json.status == 0) {
                        nameParent.addClass("error");
                        nameHelp.html(json.message).show();
                    } else {
                        window.document.location.reload();
                    }
                }
            });
        }
    });
});