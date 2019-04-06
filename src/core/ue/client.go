package ue

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func verify(IMSI string) bool {
	RAND, AUTN :=getRandAutn(IMSI)
	RES := genRES(RAND, AUTN)

	return sendRES(RES)
}

func sendRES(RES string) bool {
	resp, err := http.PostForm("http://127.0.0.1:8083/authorization",
		url.Values{"RES": {RES}},
	)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	response := string(body)
	if response == "200" {
		return true
	}

	return false
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