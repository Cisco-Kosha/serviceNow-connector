package httpclient

import (
	
	"encoding/base64"
	"bytes"
	
	"io/ioutil"
	"net/http"
	"net/url"
	// "fmt"
	// "strconv"
)
func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(username string, password string, req *http.Request, params url.Values) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.URL.RawQuery = params.Encode()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}


func GetAccount(username string, password string, params url.Values) interface{} {
	req, err := http.NewRequest("GET", "https://dev109901.service-now.com"+"/Accounts", nil)
	if err != nil {
		return nil
	}

	res := makeHttpReq(username, password, req, params)
	return res
}



func RetrieveAllTableRecords(id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("GET", "https://dev109901.service-now.com"+"api/now/table/"+id+"", nil)
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}
 
func CreateTableRecords(id string, username string, password string, body []byte, params url.Values) (string, error) { 
	req, err := http.NewRequest("POST", "https://dev109901.service-now.com"+"api/now/table/"+id+"", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func RetrieveSingleTableRecord(id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("GET", "https://dev109901.service-now.com"+"api/now/table/"+id+"/{sys_id", nil)
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}
 
func ModifyTableRecord(id string, username string, password string, body []byte, params url.Values) (string, error) { 
	req, err := http.NewRequest("PUT", "https://dev109901.service-now.com"+"api/now/table/"+id+"/{sys_id", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteTableRecord(id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", "https://dev109901.service-now.com"+"api/now/table/"+id+"/{sys_id", nil)
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}


