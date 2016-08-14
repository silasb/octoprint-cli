package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var host = "http://10.5.5.15:5000/api/"
var API_KEY = "1234"

func main() {
	println(API("files"))

	// getRequest(API("files"))

	uploadFile("test.gcode")
}

func uploadFile(path string) {
	req, err := assembleRequest(path)
	if err != nil {
		log.Panic(err)
	}

	resBody, err := callClient(req)

	if err != nil {
		log.Panic(err)
	}

	body, _ := ioutil.ReadAll(resBody)
	println(string(body))
}

func callClient(req *http.Request) (io.ReadCloser, error) {
	req.Header.Set("X-API-KEY", API_KEY)

	client := &http.Client{}
	res, err := client.Do(req)

	return res.Body, err
}

func assembleRequest(path string) (*http.Request, error) {
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

	req, err := postRequest(API("files/local"), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, err
}

func postRequest(url string, body *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)

	return req, err
}

func getRequest(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-KEY", "1234")
	res, err := client.Do(req)

	if err != nil {
		log.Panic(err)
	} else {
		defer res.Body.Close()
		_, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			log.Panic(err)
		}
	}
}

func API(resource string) string {
	return fmt.Sprintf(host+"%s", resource)
}
