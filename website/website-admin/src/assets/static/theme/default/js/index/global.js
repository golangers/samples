$(function () {
    $(window).keyup(function(event){
        if(event.keyCode == 13){
            $("#btn_login").click();
        }
    });
    $("#btn_login").bind("click", function(){
        var username = $("input[name=username]");
        var usernameHelp = username.next(".help-inline");
        var usernameParent = username.parents(".control-group");
        var password = $("input[name=password]");
        var passwordHelp = password.next(".help-inline");
        var passwordParent = password.parents(".control-group");
        var rememberStatus = ""

        username.focus(function() {
            usernameParent.removeClass("error");
            usernameHelp.hide().html("");
        });

        password.focus(function() {
            passwordParent.removeClass("error");
            passwordHelp.hide().html("");
        });

        if (username.val() == "" || password.val() == "") {
            if (username.val() == "") {
                usernameParent.addClass("error");
                usernameHelp.html("用户名不能为空").show();
            }

            if (password.val() == "") {
                passwordParent.addClass("error");
                passwordHelp.html("密码不能为空").show();
            }

            return false;
        }

        if ($("#rememberstatus").attr("checked") == "checked") {
            rememberStatus = $("#rememberstatus").val()
        }

        $.ajax({
            type : "POST",
            dataType : "json",
            data: "ajax=&username="+username.val()+"&password="+password.val()+"&rememberstatus="+rememberStatus,
            success:function(json) {
                switch (json.status) {
                    case -1: {
                        usernameParent.addClass("error");
                        usernameHelp.html(json.message).show();
                        break;
                    }
                    case 0: {
                        passwordParent.addClass("error");
                        passwordHelp.html(json.message).show();
                        break;
                    }
                    case 1: {
                        if (json.back_url != "") {
                            window.location.href = window.location.origin+json.back_url;
                        } else {
                            window.location.href = window.location.origin+"/";
                        }
                        break;
                    }
                }
            }
        });
    });
});