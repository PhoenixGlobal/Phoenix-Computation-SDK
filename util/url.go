package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// SendHttpGet send get request
func SendHttpGet(rawURL string, params url.Values, token string) (json.RawMessage, error) {
	Url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func SendHttpPost(rawURL string, rawBody json.RawMessage, token string) (json.RawMessage, error) {
	req, err := http.NewRequest("POST", rawURL, strings.NewReader(string(rawBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func SendPostMultiPart(rawURL string, body_buf *bytes.Buffer, body_writer *multipart.Writer, token string) (json.RawMessage, error) {
	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", rawURL, req_reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Connecton", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.Header.Set("token", token)
	req.ContentLength = int64(body_buf.Len())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
