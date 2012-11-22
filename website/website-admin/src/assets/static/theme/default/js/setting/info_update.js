var ck_email = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
var ck_username = /^[A-Za-z0-9_]{1,20}$/;
var ck_password  = /^[A-Za-z0-9!@#$%^&*()_]{6,20}$/;

function validateEmail(email) {
    if (!ck_email.test(email)) {
        return false;
    }
    return true;
}

function validateUsername(username) {
    if (!ck_username.test(username)) {
        return false;
    }
    return true;
}

function validatePassword(password) {
    if (!ck_password.test(password)) {
        return false;
    }
    return true;
}

$(function(){
	$("#btn_update").bind("click", function() {
		var username = $("input[name=username]").val();

		var email = $("input[name=email]");
		var emailHelp = email.next(".help-inline");
		var emailParent = email.parents(".control-group");

		var password = $("input[name=password]");
		var passwordHelp = password.next(".help-inline");
		var passwordParent = password.parents(".control-group");

		var emailVal = email.val(),
			passwordVal = password.val(),
			ajaxData = "ajax=&username=" + username,
			shouldReturn = false;

		email.focus(function() {
			emailParent.removeClass("error");
			emailHelp.hide().html("");
		});

		password.focus(function() {
			passwordParent.removeClass("error");
			passwordHelp.hide().html("");
		});

		if (emailVal == "" && passwordVal == "") {
			emailParent.addClass("error");
			emailHelp.html("要更改信息，邮箱和密码必须至少填一项").show();

			passwordParent.addClass("error");
			passwordHelp.html("邮箱和密码必须至少填一项").show();
			return;
		}

		if (emailVal != "") {
			if (validateEmail(emailVal)) {
				ajaxData += "&email=" + emailVal;
			} else {
				emailParent.addClass("error");
				emailHelp.html("请输入合法的邮箱地址！").show();
				shouldReturn = true;
			}
		}

		if (passwordVal != "") {
			if (validatePassword(passwordVal)) {
				ajaxData += "&password=" + passwordVal;
			} else {
				passwordParent.addClass("error");
				passwordHelp.html("密码必须由6到20个合法字符组成").show();
				shouldReturn = true;
			}
		}

		if (shouldReturn) {
			return;
		} else {
			$.ajax({
				url: "info_update.html",
				type: "POST",
				dataType: "json",
				data: ajaxData,
				success: function(json) {
					if (json.status == 0) {
						emailParent.addClass("error");
						emailHelp.html(json.message).show();
					} else {
						window.location.href = "/setting/info_update.html";
					}
				}
			});
		}
	});

	$('#btn_cancel').bind("click", function(){
			window.location.href = "/index.html";
	});
});