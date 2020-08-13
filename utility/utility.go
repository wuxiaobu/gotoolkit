package utility

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//CsvReadAll 读取csv
func CsvReadAll(filename string) ([]map[string]string, error) {
	records := []map[string]string{}
	file, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	cols, err := csvReader.Read()
	if err == io.EOF {
		return records, err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		row := map[string]string{}
		for index, value := range record {
			row[cols[index]] = value
		}

		records = append(records, row)
	}

	return records, nil
}

//CsvWriteLine 写入csv
func CsvWriteLine(filename string, records ...[]string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, record := range records {
		err = csvWriter.Write(record)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}

type HttpClient struct {
	http.Client
}

// 设置超时
func (h *HttpClient) SetTimeout(timeout time.Duration) {
	h.Timeout = timeout
}

// proxyUrl: http://127.0.0.1:8080
func (h *HttpClient) SetProxy(proxyUrl string) {
	if proxyUrl == "" {
		h.Transport = nil
	} else {
		urli := url.URL{}
		urlproxy, _ := urli.Parse(proxyUrl)
		transport := &http.Transport{
			Proxy:           http.ProxyURL(urlproxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
			MaxConnsPerHost: 5,
			MaxIdleConns:    5,
		}
		h.Transport = transport
	}
}

func (h *HttpClient) SendHttpRequest(requestUrl string, method string, headers *map[string]string, values *url.Values) (req *http.Request, resp *http.Response, err error) {
	method = strings.ToUpper(method)
	if values != nil {
		valueList := *values
		req, _ = http.NewRequest(method, requestUrl, bytes.NewBufferString(valueList.Encode()))
	} else {
		req, _ = http.NewRequest(method, requestUrl, nil)
	}

	if headers != nil {
		headerList := *headers
		for headerKey, headerValue := range headerList {
			req.Header.Add(headerKey, headerValue)
		}
	}

	resp, err = h.Do(req)
	return req, resp, err
}

