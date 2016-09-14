package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

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

func Run(commands []string) error {
	g := &Gcode{commands}

	body, err := json.Marshal(g)
	if err != nil {
		return err
	}

	req, err := postRequest(API("printer/command"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := callClient(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return errors.New("Bad command")
	}

	//body2, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body2))

	return nil
}

func Info() (*Printer, error) {
	req, err := getRequest(API("printer"))
	if err != nil {
		return nil, err
	}

	res, err := callClient(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Bad command")
	}

	//body2, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body2))

	var printer Printer
	err = json.NewDecoder(res.Body).Decode(&printer)
	if err != nil {
		log.Panic(err)
	}

	return &printer, nil
}

func callClient(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", Api_key)

	client := &http.Client{}
	res, err := client.Do(req)

	return res, err
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
