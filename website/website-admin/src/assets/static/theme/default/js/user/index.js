$(function() {
	$("input[op=user_delete]").click(function() {
		var h3 = $("#deleteModal .modal-header h3");
		var col = $(this).parent().siblings()
		var username = col.eq(0).text();
		h3.html("是否确认删除 " + username + " ?");
		$(this).parent().parent().attr("ready", "true");
		$("#deleteModal").modal("show");
	});

	$("#user_delete_sure").click(function() {
		var tr = $("tr[ready=true]");
		var username = tr.children().eq(0).text();

		$.ajax({
			url: "delete_user.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&username=" + username,
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

	$('input[op=user_update]').click(function() {
		var h3 = $("#updateModal .modal-header h3");
		var col = $(this).parent().siblings()
		var username = col.eq(0).text();
		h3.html("是否确认修改 " + username + " 的账户?");
		$(this).parent().parent().attr("ready", "true");
		$("#updateModal").modal("show");
	});

	$("#btn_update").bind("click", function() {
		var tr = $("tr[ready=true]");
		var username = tr.children().eq(0).text();
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
				url: "update_user.html",
				type: "POST",
				dataType: "json",
				data: ajaxData,
				success: function(json) {
					if (json.status == 0) {
						emailParent.addClass("error");
						emailHelp.html(json.message).show();
					} else {
						window.document.location.reload();
					}
				}
			});
		}
	});

	$('input[op=user_stop]').click(function() {
		var col = $(this).parent().siblings()
		var username = col.eq(0).text();
		var status = $(this).val();
		var h3 = $("#stopModal .modal-header h3");
		h3.html("是否确认" + status + " " + username + " ?");
		$("#user_stop_sure").val("确认" + status);
		$(this).parent().parent().attr("ready", "true");
		$("#stopModal").modal("show");
	});

	$("#user_stop_sure").click(function() {
		var tr = $("tr[ready=true]");
		var username = tr.children().eq(0).text();

		$.ajax({
			url: "stop_user.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&username=" + username,
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

	$('#create_user').click(function() {
		$("#createModal").modal("show");
	});

	$("#btn_create").bind("click", function() {
		var username = $("input[name=username]");
		var usernameHelp = username.next(".help-inline");
		var usernameParent = username.parents(".control-group");

		var email = $("input[name=newemail]");
		var emailHelp = email.next(".help-inline");
		var emailParent = email.parents(".control-group");

		var password = $("input[name=newpassword]");
		var passwordHelp = password.next(".help-inline");
		var passwordParent = password.parents(".control-group");

		var usernameVal = username.val(),
			passwordVal = password.val(),
			emailVal = email.val(),
			shouldReturn = false;

		username.focus(function() {
			usernameParent.removeClass("error");
			usernameHelp.hide().html("");
		});

		email.focus(function() {
			emailParent.removeClass("error");
			emailHelp.hide().html("");
		});

		password.focus(function() {
			passwordParent.removeClass("error");
			passwordHelp.hide().html("");
		});

		if (!validateUsername(usernameVal)) {
			usernameParent.addClass("error");
			usernameHelp.html("用户名必须由1到20个下字符组成：大小写字母、数字或者下划线").show();
			shouldReturn = true;
		}

		if (!validateEmail(emailVal)) {
			emailParent.addClass("error");
			emailHelp.html("请输入合法的邮箱地址").show();
			shouldReturn = true;
		}

		if (!validatePassword(passwordVal)) {
			passwordParent.addClass("error");
			passwordHelp.html("密码必须由6到20个以下字符组成：A-Za-z0-9!@#$%^&*()_").show();
			shouldReturn = true;
		}

		if (shouldReturn) {
			return;
		}

		$.ajax({
			url: "create_user.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&username=" + usernameVal + "&password=" + passwordVal + "&email=" + emailVal,
			success: function(json) {
				if (json.status == 0) {
					usernameParent.addClass("error");
					usernameHelp.html(json.message).show();
				} else {
					window.document.location.reload();
				}
			}
		});
	});
});