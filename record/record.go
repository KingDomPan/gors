package record

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Recorder struct {
	Filename string
	Command  string
	Title    string
	MaxWait  int

	screener *Screener

	once sync.Once
}

func (r *Recorder) setDefault() {
	if r.Filename == "" {
		f, err := ioutil.TempFile("/tmp", "gors")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r.Filename = f.Name()
	}

	if r.Command == "" {
		command := "bash"
		shell := os.Getenv("SHELL")
		if shell != "" {
			command = shell
		}
		r.Command = command
	}

	if r.Title == "" {
		r.Title = "Gors Terminal Recorder"
	}

	if r.MaxWait == 0 {
		r.MaxWait = 1000
	}

	if r.screener == nil {
		r.screener = NewScreener(r)
	}
}

func (r *Recorder) Execute() {
	r.once.Do(r.setDefault)

	fmt.Println("gors recording started.")
	fmt.Println("hit ctrl-d to finish.")
	fmt.Println()

	if err := r.screener.screen(r); err != nil {
		fmt.Println(err)
	}

	fmt.Println("gors recording finished.")
}
