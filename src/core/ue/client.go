package ue

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func verify(IMSI string) string {
	RAND, AUTN :=getRandAutn(IMSI)
	RES := genRES(RAND, AUTN)

	return sendRES(RES)
}

func sendRES(RES string) string {
	resp, err := http.PostForm("http://127.0.0.1:8083/authorization",
		url.Values{"RES": {RES}},
	)
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	response, err := url.ParseQuery(string(body))
	if err != nil {
		return ""
	}

	Kausd := response.Get("Kausd")

	return Kausd
}

func getRandAutn(IMSI string) (string,string) {
	resp, err := http.PostForm("http://127.0.0.1:8083/register",
		url.Values{"IMSI": {IMSI}},
	)
	if err != nil {
		return "",""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	query, err := url.ParseQuery(string(body))
	if err != nil {

	}
	RAND := query.Get("RAND")
	AUTN := query.Get("AUTN")
	_ =  resp.Body.Close()

	return RAND, AUTN
}

func genRES(RAND string, AUTN string) string {
	RES := "1234567890"

	return RES
}