1. 启动服务端

go run server.go

2. 浏览器测试

127.0.0.1:8010

3. curl测试

curl -i -H "Connection: Upgrade" -H "Upgrade: websocket" -H "Host: 127.0.0.1:8010" -H "Origin: http://127.0.0.1:8010" -H "Sec-Websocket-Key: aaaa" -H "Sec-Websocket-Version: 13" http://127.0.0.1:8010/golang

注：curl只能测试长连接的建立，无法测试后续的数据传输
