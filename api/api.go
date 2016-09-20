package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetJob() (*Job, error) {
	var job Job

	req, err := getRequest(API("job"))
	if err != nil {
		return nil, err
	}

	res, err := callClient(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func SelectFile(file File) error {
	body := `{
		"command": "select"
	}`

	resource := fmt.Sprintf("files/local/%s", file.Name)

	req, err := postRequest(API(resource), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := callClient(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return errors.New("Bad command")
	}

	//body2, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body2))

	return nil
}

func ListFiles() ([]File, error) {
	req, err := getRequest(API("files"))
	if err != nil {
		return nil, err
	}

	res, err := callClient(req)
	if err != nil {
		return nil, err
	}

	// body, _ := ioutil.ReadAll(res.Body)

	var f Files
	err = json.NewDecoder(res.Body).Decode(&f)
	if err != nil {
		return nil, err
	}

	// println(string(body))
	return f.Files, nil
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

func UploadFile(path string) (string, error) {
	req, err := assembleUploadRequest(path)
	if err != nil {
		return "error", err
	}

	res, err := callClient(req)

	if err != nil {
		return "error", err
	}

	// body, _ := ioutil.ReadAll(res.Body)

	return res.Status, nil
	// println(string(body))
}
