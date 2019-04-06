package mme

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func compareRES(RES string, XRES string) bool {
	if RES == XRES {
		return true
	}

	return false
}

func contactHSS(IMSI string, networkID string, networkType string) (string, string, string, string, error) {
	resp, err := http.PostForm("http://127.0.0.1:8084/",
		url.Values{
			"IMSI": {IMSI},
			"networkID": {networkID},
			"networkType": {networkType},
		})
	if err != nil {
		return "","","","",errors.New("fail to connect mme")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	//fmt.Println(string(body))
	query, err := url.ParseQuery(string(body))
	if err != nil {

	}
	RAND := query.Get("RAND")
	AUTN := query.Get("AUTN")
	xres := query.Get("XRES")
	kausd := query.Get("Kausd")
	_ =  resp.Body.Close()
	//fmt.Println(RAND, AUTN, xres, kausd)

	return RAND, AUTN, xres, kausd, nil
}
