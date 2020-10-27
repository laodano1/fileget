package main

var (
	tmplStr = `
<!DOCTYPE html>
<html>
<head>
    <title>Repo Sync tool</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" href="https://www.cerence.com/sites/g/files/knoqqb54601/files/favicon.ico" type="image/vnd.microsoft.icon">

    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">


</head>
<body>
<div class="container">
    <h2>Sync Tool</h2>
    <form>
        <div class="input-group">
            <input type="text" class="form-control" placeholder="Search">
            <div class="input-group-btn ">
                <button class="btn btn-default" type="submit">
                    <i class="glyphicon glyphicon-ok"></i>
                </button>
            </div>
        </div>
    </form>
    <br>
    <div class="form-group">
        <label for="comment">Result:</label>
        <div id="result" class="container"></div>
        <!--textarea class="form-control" rows="5" id="comment"></textarea-->
    </div>
</div>
    <!-- jQuery library -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <!-- Latest compiled JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    <script type="text/javascript">
        var ws = new WebSocket("ws://localhost:9999/sync");
        var cnt = 0;

        ws.onopen = function(evt) {
            //console.log("Connection open ...");
            //cnt++;
            ws.send("Hello WebSockets! ");
        };

        ws.onmessage = function(evt) {
            //console.log( "Received Message: " + evt.data);
            var logElem = document.getElementById("result");
            logElem.innerHTML += evt.data + "<br/>";
            cnt++;
            ws.send("hello server! " + cnt);
            if (cnt >= 20) {
                ws.close();
            }
        };

        ws.onclose = function(evt) {
            console.log("Connection closed.");
        };

        function WebSocketTest()
        {
            if ("WebSocket" in window)
            {
                alert("您的浏览器支持 WebSocket!");

                // 打开一个 web socket
                var ws = new WebSocket("ws://localhost:9999/sync");

                ws.onopen = function()
                {
                    // Web Socket 已连接上，使用 send() 方法发送数据
                    ws.send("发送数据");
                    alert("数据发送中...");
                };

                ws.onmessage = function (evt)
                {
                    var received_msg = evt.data;
                    alert("数据已接收..." + received_msg);
                };

                ws.onclose = function()
                {
                    // 关闭 websocket
                    alert("连接已关闭...");
                };
            }
            else
            {
                // 浏览器不支持 WebSocket
                alert("您的浏览器不支持 WebSocket!");
            }
        }
    </script>
</body>
</html>

`
)
