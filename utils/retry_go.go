package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func RetryDoRequest(method string, url string, headers map[string]string, body map[string]interface{}, typeBody string) (int, []byte) {
	var resp *http.Response
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

	timeout := time.Duration(15 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, errNewReq := http.NewRequest(method, url, reqBody)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if errNewReq != nil {
		panic(errNewReq)
	}

	err := retry.Do(
		func() error {
			var err error
			resp, err = client.Do(req)
			return err
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode, respBody
}
