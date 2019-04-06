package ue

import (
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	change = make(map[string][]byte)
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	htmlFront = `
<!DOCTYPE html>

		<html lang="zh-ch">
		<head>
		<meta charset="utf-8">
		<title>访问开始</title>
		</head>

	<body>

		<form action="/request" method="post">`
	htmlBehind = `            密钥：<input type="text" name="my"><br>
            <input type="submit" value="发 起">
        </form>

    </body>
</html>`
)

func Run(port string, path string) {
	loadHtml("register", filepath.Join(path, "register.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", register)
	mux.HandleFunc("/request", resopn)
	//mux.HandleFunc("/receive", receive)
	err := http.ListenAndServe(":" + port, mux)
	if err != nil {
		fmt.Println("bind error")

		return
	}
}

func resopn(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")
	host[1] = ":8085"
	URL := "http://"+ host[0] + host[1] + "/service"
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func register(w http.ResponseWriter, r *http.Request) {
	IMSI := r.FormValue("imsi")

	if len(IMSI) == 0 {
		_, _ = fmt.Fprintf(w, "%s", change["register"])
	} else {
		result := verify(IMSI)
		//wrong := randSeq(len(result))
		htmlMid :=  "正确密钥："+ result + ` <br>`+ ` <br><br>`
		htmlF := htmlFront + htmlMid + htmlBehind
		_, _ = w.Write([]byte(htmlF))
	}
}



func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}