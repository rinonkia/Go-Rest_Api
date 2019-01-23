# Goを使ったrest api

### 環境

 - go1.11.4
 - HTTPie/1.0.2

## setting

### 参考

[こんなに簡単! Goで作るRESTサーバー](https://qiita.com/suin/items/f32fa82d6c35a34e8d16)


httpie インストール<br>

```
$ brew install httpie
```

プログラムの実行

```
$ go run server.go
2019/XX/XX XX:XX:XX Server started.
```

別のターミナルからhttpieを使ってアクセス

```
http -v POST localhost:9999/hello "Content-Type:application/json; charset=utf-8" Name="gatapon"
```

以下のないよで帰って来れば成功

```shell
POST /hello HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 15
Content-Type: application/json; charset=utf-8
Host: localhost:9999
User-Agent: HTTPie/1.0.2

{
    "Name": "gatapon"
}

HTTP/1.1 200 OK
Content-Length: 28
Content-Type: application/json; charset=utf-8
Date: XXX
X-Powered-By: go-json-rest

{
    "Result": "Hello, gatapon"
}


```
