package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// Post request type
const (
	RequestBodyJson  string = "JSON"
	RequestBodyXForm string = "X_WWW_FORM"
)

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func GetAPI(url string, headers map[string]string) (int, map[string]interface{}) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}
	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Call API
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("API response %d", resp.StatusCode)
		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Print(scanner.Text())
		}
		return resp.StatusCode, nil
	}
	defer resp.Body.Close()

	// Parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}

	var jResult map[string]interface{}
	json.Unmarshal(body, &jResult)
	return http.StatusOK, jResult
}

func PostAPI(url string, headers map[string]string, bodyRequest map[string]interface{}) (int, Result) {
	client := &http.Client{}
	var result Result
	newBodyRequest, _ := json.Marshal(bodyRequest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(newBodyRequest))
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}
	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Call API
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		log.Errorf("Req %s, Resp %s", string(newBodyRequest), buf.String())
		return resp.StatusCode, result
	}
	defer resp.Body.Close()

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}

	result = Result{}
	json.Unmarshal(body, &result)
	return http.StatusOK, result
}

// DoRequest Create http request
func DoRequest(method string, url string, headers map[string]string, body map[string]interface{}, typeBody string) (int, []byte) {
	var err error
	var reqBody io.Reader = nil

	if headers == nil {
		headers = make(map[string]string)
	}

	if body != nil {
		if typeBody == RequestBodyJson {
			headers["Content-Type"] = "application/json"
			jBody, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}
			reqBody = bytes.NewBuffer(jBody)
		} else if typeBody == RequestBodyXForm {
			var payload string
			for key, element := range body {
				payload = payload + fmt.Sprintf("%s=%s&", key, element)
			}
			payload = strings.TrimRight(payload, "&")
			reqBody = strings.NewReader(payload)
		}
	}

	timeout := 60 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, reqBody)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode, respBody
}

func DoRequestStream(method string, url string, headers map[string]string, body map[string]interface{}, typeBody string) (int, []byte) {
	var err error
	var reqBody io.Reader = nil

	if headers == nil {
		headers = make(map[string]string)
	}

	if body != nil {
		if typeBody == RequestBodyJson {
			headers["Content-Type"] = "application/json"
			jBody, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}
			reqBody = bytes.NewBuffer(jBody)
		} else if typeBody == RequestBodyXForm {
			var payload string
			for key, element := range body {
				payload = payload + fmt.Sprintf("%s=%s&", key, element)
			}
			payload = strings.TrimRight(payload, "&")
			reqBody = strings.NewReader(payload)
		}
	}

	timeout := 60 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, reqBody)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		logrus.Error("Error reading response body:", err)
	}

	return resp.StatusCode, respBody
}
