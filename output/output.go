package output

import (
	"encoding/json"
	"io"
	"sync"
	"time"
)

type Destination struct {
	Frames   []Frame `json:"frames"`
	Version  string  `json:"version"`
	Width    int64   `json:"width"`
	Height   int64   `json:"height"`
	Duration int64   `json:"duration"`
	Command  string  `json:"command"`
	Title    string  `json:"title"`
	Term     string  `json:"term"`
	Shell    string  `json:"shell"`
}

func (d *Destination) Save(file io.ReadWriter) error {
	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	return err
}

func NewDestination(output *Output, version string, width, height int64, command, title, term, shell string) *Destination {
	return &Destination{
		Frames:   output.frames,
		Version:  version,
		Width:    width,
		Height:   height,
		Duration: output.duration,
		Command:  command,
		Title:    title,
		Term:     term,
		Shell:    shell,
	}
}

type Frame struct {
	Data  string `json:"data"`
	Delay int64  `json:"delay"`
}

type Output struct {
	frames        []Frame
	MaxWait       int64
	lastWriteTime time.Time
	sync.Mutex
	duration int64
}

func (o *Output) Write(data []byte) (int, error) {
	f := Frame{
		Data:  string(data),
		Delay: o.incrementElapsedTime(),
	}
	o.frames = append(o.frames, f)
	return len(data), nil
}

func (o *Output) incrementElapsedTime() int64 {
	o.Lock()
	defer o.Unlock()

	now := time.Now()
	delay := int64(now.Sub(o.lastWriteTime) / time.Millisecond)

	if o.MaxWait != 0 && delay > o.MaxWait {
		delay = o.MaxWait
	}
	o.duration += delay
	o.lastWriteTime = now

	return delay
}

func NewOutput(maxWait int) *Output {
	return &Output{
		frames:        make([]Frame, 0),
		MaxWait:       int64(maxWait),
		duration:      0,
		lastWriteTime: time.Now(),
	}
}
