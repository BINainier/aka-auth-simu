package hss

import (
	"fmt"
	"net/http"
	"net/url"
)

var status = true

func Run(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/F", failhandler)
	mux.HandleFunc("/R", righthandler)
	err := http.ListenAndServe(":" + port, mux)
	if err != nil {
		fmt.Println("bind error")
		return
	}
}

func failhandler(w http.ResponseWriter, r *http.Request){
	status = false
	_, _ = w.Write([]byte("已失效"))
}

func righthandler(w http.ResponseWriter, r *http.Request){
	status = true
	_, _ = w.Write([]byte("已生效"))
}

func handler(w http.ResponseWriter, r *http.Request){
	_ = r.ParseForm()
	IMSI := r.Form["IMSI"][0]
	networkID := r.Form["networkID"][0]
	networkType := r.Form["networkType"][0]
	//fmt.Println(IMSI, networkID,networkType)

	RAND, AUTN, XRES, Kausd := genXRES(IMSI, networkID, networkType)

	v := url.Values{}
	v.Set("RAND", RAND)
	v.Add("AUTN", AUTN)
	v.Add("XRES", XRES)
	v.Add("Kausd", Kausd)
	_, _ = w.Write([]byte(v.Encode()))
}

func genXRES(IMSI string, networkID string, networkType string) (string, string, string, string) {
	RAND := "1234567890"
	AUTN := "1234567890"
	var XRES string
	if status {
		XRES = "1234567890"
	} else {
		XRES = "0123456789"
	}
	Kausd := "1234567890"

	return RAND, AUTN, XRES, Kausd
}