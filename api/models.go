package api

var Host string
var Api_key string

type JobInfo struct {
	File File
}
type Job struct {
	Job JobInfo
}
type File struct {
	Name string `json:"name"`
}
type Files struct {
	Files []File
}

type Gcode struct {
	Commands []string `json:"commands"`
}

type Temp struct {
	Actual float32 `json:"actual"`
	Offset float32 `json:"offset"`
	Target float32 `json:"target"`
}
type Temperature struct {
	Bed   Temp `json:"bed"`
	Tool0 Temp `json:"tool0"`
}
type Printer struct {
	Temperature Temperature `json:"temperature"`
}
