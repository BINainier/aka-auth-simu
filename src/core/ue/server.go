package ue

import (
	"fmt"
	"io"
	"net/http"
)

var (
	change = make(map[string][]byte)
)

func init() {
	loadHtml("register", "html/register.html")
	//loadHtml("home", "html/home.html")
}

func Run(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", register)
	//mux.HandleFunc("/register", register)
	//mux.HandleFunc("/receive", receive)
	err := http.ListenAndServe(":" + port, mux)
	if err != nil {
		fmt.Println("bind error")

		return
	}
}

//func home(w http.ResponseWriter, r *http.Request) {
//	_, _ = fmt.Fprintf(w, "%s", change["home"])
//}

func register(w http.ResponseWriter, r *http.Request) {
	IMSI := r.FormValue("imsi")

	if len(IMSI) == 0 {
		_, _ = fmt.Fprintf(w, "%s", change["register"])
	} else {
		result := verify(IMSI)
		if result {
			_, _ = io.WriteString(w, "注册成功")
		} else {
			_, _ = io.WriteString(w, "注册失败")
		}
	}
}

