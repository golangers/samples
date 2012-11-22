KindEditor.ready(function(K) {
    editor = K.create('textarea[name="content"]', {
        resizeType: 1,
        allowPreviewEmoticons: false,
        allowImageUpload: false,
        themeType : 'simple',
        designMode : false,
        items: ['source', '|', 'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold', 'italic', 'underline', 'removeformat', '|', 'justifyleft', 'justifycenter', 'justifyright', 'insertorderedlist', 'insertunorderedlist', '|', 'emoticons', 'link']
    });
});

$(function() {
    $("#update #update_art_category").val(artCategory).chosen({
        allow_single_deselect: true
    });

    $("body").keydown(function(event){
        if(event.ctrlKey && event.which == 83){//ctrl + s
            //console.log("!");
            $("#btn_update").click();
            return false;
        }
    });

    $("#btn_update").bind("click", function() {
        var submit = true

        var title = $("#update_art_title");
        var titleHelp = title.next(".help-inline");
        var titleParent = title.parents(".control-group");

        var titleVal = title.val();

        title.focus(function() {
            titleParent.removeClass("error");
            titleHelp.hide().html("");
        });

        if (titleVal == "") {
            submit = false;
            titleParent.addClass("error");
            titleHelp.html("文章名不能为空").show();
        }

        var category = $("#update_art_category");
        var categoryHelp = category.next(".help-inline");
        var categoryParent = category.parents(".control-group");

        var categoryVal = category.val();

        category.focus(function() {
            categoryParent.removeClass("error");
            categoryHelp.hide().html("");
        });

        if (categoryVal == "") {
            submit = false;
            categoryParent.addClass("error");
            categoryHelp.html("文章类别不能为空").show();
        }

        editor.sync("#create_art_content");
        var content = $("#create_art_content");
        var contentHelp = content.next(".help-inline");
        var contentParent = content.parents(".control-group");

        var contentVal = content.val();

        if (contentVal == "") {
            $("#message").show().addClass("alert-error").find("span").html("文章内容为空");
        }

        if (!submit) {
            return false;
        }
    });
});