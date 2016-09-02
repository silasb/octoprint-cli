package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var Host string
var Api_key string

type JobInfo struct {
	File File
}
type Job struct {
	Job JobInfo
}

func GetJob() Job {
	var j Job

	req, err := getRequest(API("job"))
	if err != nil {
		log.Panic(err)
	}

	res, err := callClient(req)
	if err != nil {
		log.Panic(err)
	}

	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		log.Panic(err)
	}

	return j
}

type File struct {
	Name string `json:"name"`
}
type Files struct {
	Files []File
}

func ListFiles() []File {
	req, err := getRequest(API("files"))
	if err != nil {
		log.Panic(err)
	}

	res, err := callClient(req)
	if err != nil {
		log.Panic(err)
	}

	// body, _ := ioutil.ReadAll(res.Body)

	var f Files
	err = json.NewDecoder(res.Body).Decode(&f)
	if err != nil {
		log.Panic(err)
	}

	// println(string(body))
	return f.Files
}

func callClient(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", Api_key)

	client := &http.Client{}
	res, err := client.Do(req)

	return res, err
}

func assembleUploadRequest(path string) (*http.Request, error) {
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

func getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	return req, err
}

func API(resource string) string {
	return fmt.Sprintf(Host+"%s", resource)
}

func UploadFile(path string) string {
	req, err := assembleUploadRequest(path)
	if err != nil {
		log.Panic(err)
	}

	res, err := callClient(req)

	if err != nil {
		log.Panic(err)
	}

	// body, _ := ioutil.ReadAll(res.Body)

	return res.Status
	// println(string(body))
}
