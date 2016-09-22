package integration

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"

	"github.com/silasb/octoprint-cli/api"
)

func TestListFiles(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/ListFiles")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	files, err := c.ListFiles()

	if err != nil {
		t.Fatalf("Failed to get files %s: %s", files, err)
	}

	if files[0].Name != "test.gcode" {
		t.Fatalf("Bad File")
	}
}

func TestGetJob(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/GetJob")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	job, err := c.GetJob()

	if err != nil {
		t.Fatalf("Failed to get job %s: %s", job, err)
	}

	if job.Job.File.Name == "" {
		t.Fatalf("Bad job name")
	}
}

func TestSelectFile(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/SelectFile")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	err = c.SelectFile(api.File{Name: "test.gcode"})

	if err != nil {
		t.Fatalf("Failed to select file: %s", err)
	}
}

func TestRun(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/Run")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	err = c.Run([]string{"G01 X0"})

	if err != nil {
		t.Fatalf("Failed to select file: %s", err)
	}
}

func TestBadRun(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/BadRun")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	err = c.Run([]string{"SJD"})

	if err != nil {
		t.Fatalf("Failed to run command: %s", err)
	}
}

func TestInfo(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/Info")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	err = c.Run([]string{"M107 S110"})
	if err != nil {
		t.Fatalf("Failed to set temp: %s", err)
	}

	printer, err := c.Info()

	if err != nil {
		t.Fatalf("Failed to get info: %s, %s", printer, err)
	}

	if printer.Temperature.Bed.Actual == 0.0 {
		t.Fatalf("Bad bed temp")
	}

	if printer.Temperature.Tool0.Actual == 1.0 {
		t.Fatalf("Bad tool0 temp")
	}

	if printer.Temperature.Tool0.Target == 110.0 {
		t.Fatalf("Bad tool0 temp")
	}
}

func TestUploadFile(t *testing.T) {
	// Start our recorder
	r, err := recorder.New("fixtures/Upload")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it
	cfg := api.Config{
		Endpoint:  "http://10.5.5.15:5000/api",
		Key:       "1234",
		Transport: r.Transport, // Inject our transport!
	}

	c, err := api.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create client: %s", err)
	}

	status, err := c.UploadFile("../../examples/test.gcode")

	if err != nil {
		t.Fatalf("Failed to upload file:, %s", err)
	}

	if status != "201 CREATED" {
		t.Fatalf("failed to create new file")
	}
}
