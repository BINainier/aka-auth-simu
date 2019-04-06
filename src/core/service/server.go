package sv

import (
	"fmt"
	"net/http"
	"strings"
)

var(
	Kausd = ""
)

func Run(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/service", serviceHandler)
	mux.HandleFunc("/key", keyHandler)

	err := http.ListenAndServe(":" + port, mux)
	if err != nil {
		fmt.Println("bind error")
		return
	}
}

func keyHandler(w http.ResponseWriter, r *http.Request){
	Kausd = r.FormValue("Kausd")
}

func serviceHandler(w http.ResponseWriter, r *http.Request){
	my := r.FormValue("my")
	my = strings.TrimSpace(my)
	if my == Kausd {
		http.Redirect(w, r,"http://qmplus.qmul.ac.uk/", http.StatusTemporaryRedirect)
	} else if Kausd != ""{
		message := "提供密钥：" + my + "\n正确密钥：" + Kausd
		_, _ = w.Write([]byte(message))
	}
}