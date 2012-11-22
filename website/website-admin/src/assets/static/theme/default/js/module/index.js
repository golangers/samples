$(function() {
	$("input[op=module_delete]").click(function() {
		var col = $(this).parent().siblings()
		var modulename = col.eq(0).text();
		var modulepath = col.eq(1).text();
		var h3 = $("#deleteModal .modal-header h3");
		h3.html("是否确认删除 " + modulename + "(" + modulepath + ") ?");
		$(this).parent().parent().attr("ready", "true");
		$("#deleteModal").modal("show");
	});

	$("#module_delete_sure").click(function() {
		var tr = $("tr[ready=true]");
		var modulename = tr.children().eq(0).text();
		var modulepath = tr.children().eq(1).text();
		$.ajax({
			url: "delete_module.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&modulename=" + modulename + "&modulepath=" + modulepath,
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

	$('input[op=module_stop]').click(function() {
		var col = $(this).parent().siblings()
		var modulename = col.eq(0).text();
		var modulepath = col.eq(1).text();
		var status = $(this).val();
		var h3 = $("#stopModal .modal-header h3");
		h3.html("是否确认" + status + " " + modulename + " ?");
		$("#module_stop_sure").val("确认" + status);
		$(this).parent().parent().attr("ready", "true");
		$("#stopModal").modal("show");

	});

	$("#module_stop_sure").click(function() {
		var tr = $("tr[ready=true]");
		var modulepath = tr.children().eq(1).text();

		$.ajax({
			url: "stop_module.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&modulepath=" + modulepath,
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
		var modulename = $("input[name=modulename]");
		var modulenameHelp = modulename.next(".help-inline");
		var modulenameParent = modulename.parents(".control-group");
		var modulepath = $("input[name=modulepath]");
		var modulepathHelp = modulepath.next(".help-inline");
		var modulepathParent = modulepath.parents(".control-group");

		var modulenameVal = modulename.val();
		var modulepathVal = modulepath.val();

		modulename.focus(function() {
			modulenameParent.removeClass("error");
			modulenameHelp.hide().html("");
		});

		modulepath.focus(function() {
			modulenameParent.removeClass("error");
			modulenameHelp.hide().html("");
		});

		if (modulenameVal == "") {
			modulenameParent.addClass("error");
			modulenameHelp.html("模块名不能为空").show();
		}

		if (modulenameVal == "") {
			modulepathParent.addClass("error");
			modulepathHelp.html("模块路径不能为空").show();
		}

		$.ajax({
			url: "create_module.html",
			type: "POST",
			dataType: "json",
			data: "ajax=&modulename=" + modulenameVal + "&modulepath=" + modulepathVal,
			success: function(json) {
				console.log(json);
				if (json.status == 0) {
					modulenameParent.addClass("error");
					modulenameHelp.html(json.message).show();
				} else if (json.status == 1) {
					window.document.location.reload();
				}
			}
		});
	});
});