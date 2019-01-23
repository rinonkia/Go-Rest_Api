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

## 内容

叩いたコード<br>
`http -v POST localhost:9999/hello "Content-Type:application/json; charset=utf-8" Name="gatapon"`

「`localhost:9999/hello`にPOSTメソッドでJSON形式の`Name="gatapon"`を渡す。」

```golang
// 第一引数 "Content-Type:application/json; charset=utf-8"
// 第二引数 Name="gatapon"
func postHello(w rest.ResponseWriter, req *rest.Request) {
	input := postHelloInput{}
    err := req.DecodeJsonPayload(&input)

	if err != nil { // DecodeJsonPayload で検出されたエラーを返す
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
```

[DecodeJsonPayload](https://github.com/ant0ine/go-json-rest/blob/ebb33769ae013bd5f518a8bac348c310dea768b8/rest/request.go#L34) 

`DecodeJsonPayload`で様々なバリデーションを行なっており、当てはまった場合はエラーを返す。<br>

 - 文字列の終端まで読み込み、バイトスライス
 - 文字列が0の時のバリデーション
 - Json形式の文字列をパース：json.Unmarshal

