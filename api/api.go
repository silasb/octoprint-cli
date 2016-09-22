package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	client *http.Client
	cfg    Config
}

func (c *Client) API(resource string) string {
	return fmt.Sprintf(c.cfg.Endpoint+"/%s", resource)
}

func New(cfg Config) (*Client, error) {
	return newClient(&cfg)
}

func newClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	client := &Client{
		client: &http.Client{
			Transport: cfg.Transport,
		},
		cfg: *cfg,
	}

	return client, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", c.cfg.Key)

	return c.client.Do(req)
}

/* */

func (c *Client) GetJob() (*Job, error) {
	var job Job

	req, err := getRequest(c.API("job"))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (c *Client) SelectFile(file File) error {
	body := `{
		"command": "select"
	}`

	resource := fmt.Sprintf("files/local/%s", file.Name)

	req, err := postRequest(c.API(resource), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Do(req)
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

func (c *Client) ListFiles() ([]File, error) {
	req, err := getRequest(c.API("files"))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
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

func (c *Client) Run(commands []string) error {
	g := &Gcode{commands}

	body, err := json.Marshal(g)
	if err != nil {
		return err
	}

	req, err := postRequest(c.API("printer/command"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Do(req)
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

func (c *Client) Info() (*Printer, error) {
	req, err := getRequest(c.API("printer"))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
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

func (c *Client) UploadFile(path string) (string, error) {
	req, err := c.assembleUploadRequest(path)
	if err != nil {
		return "error", err
	}

	res, err := c.Do(req)

	if err != nil {
		return "error", err
	}

	// body, _ := ioutil.ReadAll(res.Body)

	return res.Status, nil
	// println(string(body))
}
