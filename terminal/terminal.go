package terminal

import (
	"bufio"
	"io"
	"os"
	"sync"
)

type CmdCallback func(cmd []byte) error

type cmdTerminal struct {
	reader    io.ReadCloser
	bufReader *bufio.Reader
	writer    io.WriteCloser
	sync.Mutex
	cmdCallbacks []CmdCallback
}

func (ct *cmdTerminal) OnCmdCallback(cb CmdCallback) {
	ct.cmdCallbacks = append(ct.cmdCallbacks, cb)
}

func (ct *cmdTerminal) ReadBytes(delim byte) ([]byte, error) {
	return ct.bufReader.ReadBytes(delim)
}

func (ct *cmdTerminal) Write(p []byte) (int, error) {
	return ct.writer.Write(p)
}

func (ct *cmdTerminal) Close() error {
	if ct.reader != nil {
		if err := ct.reader.Close(); err != nil {
			return err
		}

	}
	if ct.writer != nil {
		if err := ct.writer.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (ct *cmdTerminal) IOSelect(errors chan error) {
	for {
		cmd, err := ct.ReadBytes('\r')
		if err != nil {
			return
		}
		cmd = cmd[:len(cmd)-1]
		for _, callback := range ct.cmdCallbacks {
			err := callback(cmd)
			if err != nil {
				errors <- err
				return
			}
		}
	}
}

func NewCmdTermial() (*cmdTerminal, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	return &cmdTerminal{
		reader:       r,
		writer:       w,
		bufReader:    bufio.NewReader(r),
		cmdCallbacks: make([]CmdCallback, 0),
	}, nil
}
