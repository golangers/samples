var editor;
KindEditor.ready(function(K) {
    editor = K.create('textarea[name="content"]', {
        resizeType: 1,
        allowPreviewEmoticons: false,
        allowImageUpload: false,
        designMode : false,
        items: ['source', '|', 'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold', 'italic', 'underline', 'removeformat', '|', 'justifyleft', 'justifycenter', 'justifyright', 'insertorderedlist', 'insertunorderedlist', '|', 'emoticons', 'link']
    });
});



$(function() {

    var all_cols = $("#art_cols > tr");
    function col_with_cat(category){
        var cols = [];
        all_cols.each(function(){
            if($(this).attr("category")==category){
                cols.push($(this));
            }
        });
        return cols;
    }
    function col_of_all(){
        var cols = [];
        all_cols.each(function(){
            cols.push($(this));
        });
        return cols;
    }

    var curr_cols = col_of_all();
    var num_in_page = 20;
    var curr_index, num, page_num;

    function init(){
        $("#next_page").removeClass("disabled");

        curr_index = 1;
        num = curr_cols.length;
        page_num = Math.ceil(num / num_in_page);
        $("#prev_page").addClass("disabled");
        if(page_num==1){
           $("#next_page").addClass("disabled");
        }
        $("#page_num").html(page_num);
        $("#_pages").html("");
        for(i=1;i<=page_num;i++){
            li = '<li><input class="btn" type="button" id="'+i+'_page" value="'+i+'"></li>';
            $("#_pages").append(li);
        }
        $("#_pages > li:nth-child("+curr_index+") > input").addClass("disabled");
    }

    function show_page(){
        var min = (curr_index-1) * num_in_page;
        var max = curr_index * num_in_page;
        all_cols.each(function(){
            $(this).hide();
        });
        for(var i=min;i<num && i<max;i++){
            curr_cols[i].show();
        }

        if(curr_index==1){
            $("#prev_page").addClass("disabled");
        }
        if(curr_index==page_num){
            $("#next_page").addClass("disabled");
        }
    }

    function switch_page(new_index){
        if(new_index>=1 && new_index<=page_num && new_index!=curr_index){
            $("#prev_page").removeClass("disabled");
            $("#next_page").removeClass("disabled");

            $("#_pages > li:nth-child("+curr_index+") > input").removeClass("disabled");
            curr_index = new_index;
            show_page();
            $("#_pages > li:nth-child("+curr_index+") > input").addClass("disabled");
        }
    }

    init();
    show_page();

    $("#category_select").change(function(){
        if(this.value == 0){
           curr_cols = col_of_all(); 
        }else{
            category_id = this.value;
            curr_cols = col_with_cat(category_id);
        }
        init();
        show_page();
    });

    $("#prev_page").click(function(){
        switch_page(curr_index-1);
    });
    $("#next_page").click(function(){
        switch_page(curr_index+1);
    });
    $("#_pages > li > input[type=button]").live('click', function(){
        switch_page(parseInt(this.value));
    });



    $("#create #create_art_category").chosen({
        allow_single_deselect: true
    });
    $("#category_select").chosen({
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

    $("#btn_create").bind("click", function() {
        var submit = true

        var title = $("#create_art_title");
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

        var top = $("#create_art_top");
        var topVal = top.val();
        var hot = $("#create_art_hot");
        var hotVal = hot.val();
        var order = $("#create_art_order");
        var orderVal = order.val();
        var target_href = $("#create_art_target_href");
        var target_hrefVal = target_href.val();

        var category = $("#create_art_category");
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

        var summary = $("#create_art_summary");

        var summaryVal = summary.val();

        if (contentVal == "") {
            alert("文章内容为空");
        }

        if (submit) {
            $.ajax({
                url: "create.html",
                type: "POST",
                dataType: "json",
                data: "ajax=&title=" + titleVal + "&top=" + topVal + "&hot=" + hotVal + "&order=" + orderVal + "&target_href=" + target_hrefVal + "&category=" + categoryVal + "&content=" + encodeURIComponent(contentVal) + "&summary=" + encodeURIComponent(summaryVal),
                success: function(json) {
                    if (json.status == 0) {
                        $("#message").show().addClass("alert-error").find("span").html(json.message);
                    } else if (json.status == 1) {
                        window.document.location.reload();
                    }
                }
            });
        }
    });
});