package main

import (
	"log"
	"net/http"

	// githubから直接パッケージをいっポートできて便利
	"github.com/ant0ine/go-json-rest/rest"
)

// 入力の定義
type postHelloInput struct {
	Name string
}

// 出力の定義
type postHelloOutput struct {
	Result string
}

func postHello(w rest.ResponseWriter, req *rest.Request) {
	input := postHelloInput{}
	// 各種バリデーションが行われている README.md参照
	err := req.DecodeJsonPayload(&input)

	// バリデーションで引っかかったエラーを返す
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if input.Name == "" {
		rest.Error(w, "Name is required", 400)
	}

	log.Printf("%#v", input)

	w.WriteJson(&postHelloOutput{
		"Hello, " + input.Name,
	})
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/hello", postHello),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started.")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}
