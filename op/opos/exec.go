package opos

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"unsafe"
)

// os exec command eg: bash/sh

// Cmder Cmder
type Cmder struct {
	terminal []string
}

func (c *Cmder) getCmd(arg ...string) *exec.Cmd {
	if c.terminal == nil || len(c.terminal) < 2 {
		c.terminal = make([]string, 2)
		if runtime.GOOS == "windows" {
			c.terminal[0] = "cmd"
			c.terminal[1] = "/c"
		} else {
			c.terminal[0] = "/bin/sh"
			c.terminal[1] = "-c"
		}
	}
	var cmdArg []string
	cmdArg = append(cmdArg, c.terminal...)
	cmdArg = append(cmdArg, arg...)
	return exec.Command(cmdArg[0], cmdArg[1:]...)
}

// Run Run
func (c *Cmder) Run(cmdLine string, timeout ...int) *Result {
	cmd := c.getCmd(cmdLine)
	ret := new(Result)

	cmd.Stdout = &ret.buf
	cmd.Stderr = &ret.buf
	cmd.Env = os.Environ()

	ret.err = cmd.Start()
	if ret.err != nil {
		return ret
	}

	if len(timeout) == 0 || timeout[0] <= 0 {
		ret.err = cmd.Wait()
		return ret
	}

	timer := time.NewTimer(time.Duration(timeout[0]) * time.Second)
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case ret.err = <-done:
		timer.Stop()
	case <-timer.C:
		if err := cmd.Process.Kill(); err != nil {
			ret.err = fmt.Errorf("command timed out and killing process fail: %s", err.Error())
		} else {
			<-done
			ret.err = errors.New("command timed out")
		}
	}
	return ret
}

// Result Result
type Result struct {
	buf bytes.Buffer
	err error
	str *string
}

func (r *Result) Buf() bytes.Buffer {
	return r.buf
}

// Err Err
func (r *Result) Err() error {
	if r.err == nil {
		return nil
	}
	r.err = errors.New(r.String())
	return r.err
}

// String String
func (r *Result) String() string {
	if r.str == nil {
		b := bytes.TrimSpace(r.buf.Bytes())
		if r.err != nil {
			b = append(b, ' ', '(')
			b = append(b, r.err.Error()...)
			b = append(b, ')')
		}
		r.str = (*string)(unsafe.Pointer(&b))
	}
	return *r.str
}
