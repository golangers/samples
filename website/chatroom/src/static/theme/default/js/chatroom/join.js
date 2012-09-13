$(function() {
    var conn;

    $.getDocHeight = function() {
        return Math.max(
        $(document).height(), $(window).height(), /* For opera: */
        document.documentElement.clientHeight);
    };

    function chatResized() {
        var doc_width = document.body.clientWidth;
        var chat_width = document.body.clientWidth - 405;
        var chat_height = $.getDocHeight() - 105;
        if (chat_width > 0) {
            $('#chat-column, #input-box').css('width', chat_width);
            $("#appendedPrependedInput").css('width', (chat_width - 80));
        }
        $('#chat-column').css('height', chat_height);

    };

    chatResized();
    window.onresize = chatResized;

    function addMessage(textMessage) {

        $("#msg-template .userpic").html("<img src='" + textMessage.UserInfo.Gravatar + "'>")
        $("#msg-template .msg-time").html(textMessage.Time);
        $("#msg-template .user").html(textMessage.UserInfo.Name + ":");
        $("#msg-template .content").html(textMessage.Content);
        $("#chat-messages").append($("#msg-template").html());
        $('#chat-column')[0].scrollTop = $('#chat-column')[0].scrollHeight;
    };

    function updateUsers(userStatus) {
        var users = userStatus.Users;
        $("#user-list").html("");
        $("#user-list").append("<li class='nav-header'>Who is here?</li>");
        for (var v = 0; v < users.length; v++) {
            $("#user-list").append("<li class='user-online'> <img src='" + users[v].Gravatar + "'>" + users[v].Name + "</li>")
        }
    };

    function errorMessage(msg) {
        $("#msg-template .content").html(msg);
        $("#chat-messages").append($("#msg-template").html());
        $('#chat-column')[0].scrollTop = $('#chat-column')[0].scrollHeight;
    };

    $("#msg_form").submit(function() {
        var msg = $("#appendedPrependedInput");
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        conn.send(msg.val());
        msg.val("");
        return false
    });

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://"+WebSocketHost+"/chat?email="+Email);
        conn.onopen = function() {};

        conn.onmessage = function(evt) {
            var data = JSON.parse(evt.data);
            switch (data.MType) {
            case "text_mtype":
                addMessage(data.TextMessage)
                break;
            case "status_mtype":
                updateUsers(data.UserStatus)
                break;
            default:
            }
        };

        conn.onerror = function() {
            errorMessage("<strong> An error just occured.<strong>")
        };

        conn.onclose = function() {
            errorMessage("<strong>Connection closed.<strong>")
        };
    } else {
        errorMessage("Your browser does not support WebSockets.");
    }
});
