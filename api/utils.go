package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func postRequest(url string, body *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)

	return req, err
}

func getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	return req, err
}

func (c *Client) assembleUploadRequest(path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	// if we need additional params to be passed in.
	// for key, val := range params {
	// 	_ = writer.WriteField(key, val)
	// }

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := postRequest(c.API("files/local"), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, err
}
