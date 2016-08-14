package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

var host = "http://10.5.5.15:5000/api/"
var API_KEY = "1234"

// var (
// 	uploadCommand = flag.NewFlagSet("upload", flag.ExitOnError)
// 	questionFlag  = uploadCommand.String("file", "", "Question that you are asking for")
// )

func init() {

	// if len(os.Args) == 1 {
	// 	fmt.Println("usage: siri <command> [<args>]")
	// 	fmt.Println("The most commonly used git commands are: ")
	// 	fmt.Println(" ask   Ask questions")
	// 	fmt.Println(" send  Send messages to your contacts")
	// 	os.Exit(2)
	// }

	// switch os.Args[1] {
	// case "upload":
	//     os.Args[2:]
	// 	uploadCommand.Parse()
	// default:
	// 	fmt.Printf("%q is not valid command.\n", os.Args[1])
	// 	os.Exit(2)
	// }

	// if uploadCommand.Parsed() {
	// 	if *questionFlag == "" {
	// 		fmt.Println("Please supply the question using -question option.")
	// 		os.Exit(2)
	// 	}
	// 	fmt.Printf("You asked: %q\n", *questionFlag)
	// }

	// username := flag.String("user", "root", "Username for this server")
	// flag.Parse()
	// fmt.Printf("Your username is %q.", *username)

}

func main() {
	app := cli.NewApp()
	app.Name = "octoprint"
	app.Usage = ""
	app.Commands = []cli.Command{
		{
			Name:      "upload",
			Aliases:   []string{"u"},
			Usage:     "upload files",
			ArgsUsage: "[files]",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					for _, file := range c.Args() {
						fmt.Print("Uploading file: ", file)
						status := uploadFile(file)
						fmt.Println(" =>", status)
					}
					return nil
				}

				return errors.New("missing file")
			},
		},
		{
			Name:    "files",
			Aliases: []string{"f"},
			Usage:   "list files",
			Action: func(c *cli.Context) error {
				job := getJob()
				files := listFiles()
				for _, file := range files {
					fmt.Print(file.Name)
					if job.Job.File.Name == file.Name {
						fmt.Println(" *")
					} else {
						println()
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func uploadFile(path string) string {
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

type JobInfo struct {
	File File
}
type Job struct {
	Job JobInfo
}

func getJob() Job {
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

func listFiles() []File {
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
	req.Header.Set("X-API-KEY", API_KEY)

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
	return fmt.Sprintf(host+"%s", resource)
}
