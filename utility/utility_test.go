package utility

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func TestCsv(t *testing.T) {
	filename := "../test/sample.csv"

	record := []string{"id", "name"}
	err := CsvWriteLine(filename, record)
	if err != nil {
		t.Error(err)
	}

	_, err = CsvReadAll(filename)
	if err != nil {
		t.Error(err)
	}

	//os.Remove(filename)
}

func TestSendHttpRequestGet(t *testing.T) {
	pageUrl := "https://www.google.com"
	client := &HttpClient{}
	client.SetTimeout(5 * time.Second)
	client.SetProxy("http://127.0.0.1:7890")
	req, resp, err := client.SendHttpRequest(pageUrl, "GET", nil, nil)

	fmt.Println(req)
	fmt.Printf("%#v", resp)

	if err != nil {
		t.Error(err)
	}
}

func TestSendHttpRequestPost(t *testing.T) {
	pageUrl := "https://www.bing.com"
	client := &HttpClient{}
	values := &url.Values{}
	values.Add("Key", "Value")
	_, _, err := client.SendHttpRequest(pageUrl, "POST", nil, values)
	if err != nil {
		t.Error(err)
	}
}