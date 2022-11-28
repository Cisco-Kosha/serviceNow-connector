package httpclient

import (
	"bytes"
	"encoding/base64"
	"fmt"

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

// func GetAccount(username string, password string, params url.Values) interface{} {
// 	req, err := http.NewRequest("GET", "https://dev109901.service-now.com"+"/Accounts", nil)
// 	if err != nil {
// 		return nil
// 	}

// 	res := makeHttpReq(username, password, req, params)
// 	return res
// }

func RetrieveAllTableRecords(tableName string, url string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("GET", url+"api/now/table/"+tableName, nil)
	fmt.Println("Retrieve all table records")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func CreateTableRecords(tableName string, url string, username string, password string, body []byte, params url.Values) (string, error) {
	req, err := http.NewRequest("POST", url+"api/now/table/"+tableName, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func RetrieveSingleTableRecord(tableName string, sysId string, url string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("GET", url+"api/now/table/"+tableName+"/"+sysId, nil)
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func ModifyTableRecord(tableName string, sysId string, url string, username string, password string, body []byte, params url.Values) (string, error) {
	req, err := http.NewRequest("PUT", url+"api/now/table/"+tableName+"/"+sysId, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteTableRecord(tableName string, sysId string, url string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"api/now/table/"+tableName+"/"+sysId, nil)
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}
