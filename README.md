# GolangでRESTfulなAPIサーバー作ってみた

## 概要

GolangでRESTサーバを起動し、POSTメソッドを受け付けるAPIを作成しました。（用語の使い方あってる？）<br>
`go-json-rest`パッケージを利用して、JSON形式のデータをやりとりする。<br>

 1. httpieを使って `localhost:9999/hello`にアクセス
 2. JSON形式の`Name="gatapon"`のデータを渡す
 3. `Hello, gatapon`のレスポンスを受け取る

### 環境

 - go1.11.4
 - HTTPie/1.0.2
 
### 参考

 - [こんなに簡単! Goで作るRESTサーバー](https://qiita.com/suin/items/f32fa82d6c35a34e8d16)
 - https://golang.org/pkg/
 - https://github.com/ant0ine/go-json-rest#api-and-static-files


## setting

検証にhttpieを利用するのでインストール<br>

```
$ brew install httpie
```

適当なフォルダを作り、`git clone`<br>

```
git clone https://github.com/rinonkia/Go-Rest_Api.git
```

プログラムの実行。これでサーバが起動しAPIにアクセスができます。<br>

```
$ go run server.go
2019/XX/XX XX:XX:XX Server started.
```

別のターミナルからhttpieを使ってアクセス<br>

```
$ http -v POST localhost:9999/hello "Content-Type:application/json; charset=utf-8" Name="gatapon"
```

以下の内容でレスポンスが返ってくれば成功<br>

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
    "Result": "Hello, gatapon"
}
```

## 詳細（作成中）

ターミナルで叩いたコマンドについて振り返ります。<br>

```
http -v POST localhost:9999/hello "Content-Type:application/json; charset=utf-8" Name="gatapon"
```

これを簡単に訳すると、以下のようなことをしています。<br>

「`localhost:9999/hello`にPOSTメソッドでJSON形式の`Name="gatapon"`を渡す。」<br>

このコマンドを受け取ったサーバーの動きを見ていきます。

### postHello()

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
 - 文字列が0の時のバリデーション
 - Json形式の文字列をパース：json.Unmarshal

### 　api := rest.NewApi()

[NewApi()](https://github.com/ant0ine/go-json-rest/blob/ebb33769ae013bd5f518a8bac348c310dea768b8/rest/api.go#L14)

`NewApi()`は3つの強力なメソッドを持っており、それらを使うことで艱難にAPIサーバーを立ち上げることができます。<br>

```golang

api := NewApi()

api.Use()
api.SetApp()
api.MakeHandler()
```

 - api.Use => APIサーバーの環境に合わせて用意されたMiddlewareの設定ができる
   - DefaultDevStack    : 開発時に利用するMiddleware群
   - DefaultProdStack   : プロダクト時に利用するMiddleware群
   - DefaultCommonStack : どちらにしても一般的に利用するMiddleware群
 - api.SetApp => 主にrest.MakeRouterのrouterをセットするメソッド
 - api.MakeHandler() 


## 感想

 go-json-restを使えばとても簡単にRESTfullなAPIサーバーが作成できる。<br>
 これを使えば、思ったより早く簡単なアプリが作成できそう。<br>
