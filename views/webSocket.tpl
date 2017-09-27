<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>

    <script>
        $(function () {
            var ws = new WebSocket('ws://' + window.location.host + '/ws');
            ws.onmessage = function (e) {
                $('<li>').text(event.data).appendTo($ul);
            };
            setInterval(function () {
                ws.send(JSON.stringify({
                    message: "test"
                }));
            }, 3000);
            var $ul = $('#msg-list');
        });
    </script>
</head>
<body>
<ul id="msg-list"></ul>
</body>
</html>