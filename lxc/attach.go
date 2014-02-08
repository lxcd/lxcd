package lxc

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	"github.com/caglar10ur/lxc"
)

type AttachArgs struct {
	Name    string
	Command []string
}

type AttachReply struct {
	StatusCode int
	Stdout     string
	Stderr     string
}

func (l *LXC) Attach(args *AttachArgs, reply *AttachReply) error {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	var wg sync.WaitGroup

	c, err := lxc.NewContainer(args.Name, lxc.DefaultConfigPath())
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer lxc.PutContainer(c)

	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		log.Println("Pipe Error: %s\n", err.Error())
		return err
	}
	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		log.Println("Pipe Error: %s\n", err.Error())
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(&stdoutBuf, stdoutReader)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(&stderrBuf, stderrReader)
	}()
	if err := c.RunCommand(0, stdoutWriter.Fd(), stderrWriter.Fd(), args.Command...); err != nil {
		log.Println("ERROR: %s\n", err.Error())
		return err
	}
	stdoutWriter.Close()
	stderrWriter.Close()
	stdoutReader.Close()
	stderrReader.Close()
	reply.Stdout = stdoutBuf.String()
	reply.Stderr = stderrBuf.String()
	wg.Wait()
	return nil
}
