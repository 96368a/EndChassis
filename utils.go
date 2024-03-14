package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"log"
	"os/exec"
	"sync"
)

type Utils struct {
	ctx context.Context
}

func (u Utils) ExecCommand(command string, index int) error {
	log.Print("ExecCommand:", command)
	c := exec.Command("cmd", "/C", command) // windows
	//c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				runtime.EventsEmit(*wailsContext, "setCommandStatus", index, false)
				return
			}
			byte2String := ConvertByte2String([]byte(readString), "GB18030")
			runtime.EventsEmit(*wailsContext, "addCommandResult", index, byte2String)
			fmt.Print(byte2String)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			byte2String := ConvertByte2String([]byte(readString), "GB18030")
			fmt.Print(byte2String)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)
