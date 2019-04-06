package mme

import (
	"fmt"
	"net/http"
	"net/url"
)

var(
	networkID = "1"
	networkType = "2"
	XRES string
	Kausd string
)

func Run(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/authorization", authorizationHandler)
	err := http.ListenAndServe(":" + port, mux)
	if err != nil {
		fmt.Println("bind error")
		return
	}
}

func authorizationHandler(w http.ResponseWriter, r *http.Request){
	_ = r.ParseForm()
	RES := r.Form["RES"][0]

	if compareRES(RES, XRES) {
		authInfo := []byte("200")
		_, _ = w.Write(authInfo)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	IMSI, found := r.Form["IMSI"]
	if !found {
		return
	}
	RAND, AUTN, xres, kausd, err := contactHSS(IMSI[0], networkID, networkType)
	XRES = xres
	Kausd = kausd
	if err != nil {
		return
	}

	v := url.Values{}
	v.Set("RAND", RAND)
	v.Add("AUTN", AUTN)
	_, _ = w.Write([]byte(v.Encode()))
}

