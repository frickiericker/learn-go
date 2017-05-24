package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"

	"github.com/kr/pty"
	"github.com/pkg/errors"
)

const OutputName = "output.dat"
const CommandName = "top"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	recorder := new(Recorder)
	defer recorder.Close()
	return recorder.Run()
}

type Recorder struct {
	output   *os.File
	command  *exec.Cmd
	terminal *os.File
	err      error
}

func (rec *Recorder) Close() {
	if rec.terminal != nil {
		rec.terminal.Close()
	}
	if rec.command != nil {
		rec.command.Process.Signal(os.Interrupt)
	}
	if rec.output != nil {
		rec.output.Close()
	}
}

func (rec *Recorder) Run() error {
	rec.createOutput()
	rec.createCommand()
	return rec.runLoop()
}

func (rec *Recorder) createOutput() {
	if rec.err != nil {
		return
	}

	output, err := os.Create(OutputName)
	if err != nil {
		rec.err = errors.Wrap(err, "creating record file")
	}
	rec.output = output
}

func (rec *Recorder) createCommand() {
	if rec.err != nil {
		return
	}

	command := exec.Command(CommandName)
	terminal, err := pty.Start(command)
	if err != nil {
		rec.err = errors.Wrap(err, "creating pty")
		return
	}
	if err := setTerminalSize(terminal.Fd(), 80, 24); err != nil {
		rec.err = errors.Wrap(err, "setting terminal size")
		return
	}
	rec.command = command
	rec.terminal = terminal
}

func (rec *Recorder) runLoop() error {
	if rec.err != nil {
		return rec.err
	}

	dataChan := make(chan []byte)
	errorChan := make(chan error)
	go func() {
		for {
			buffer := make([]byte, 1024)
			size, err := rec.terminal.Read(buffer)
			if err != nil {
				if err != io.EOF {
					errorChan <- errors.Wrap(err, "reading from pty")
				}
				break
			}
			dataChan <- buffer[:size]
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	for {
		select {
		case data := <-dataChan:
			if _, err := rec.output.Write(data); err != nil {
				return errors.Wrap(err, "writing to record file")
			}
		case err := <-errorChan:
			return err
		case <-signalChan:
			return nil
		}
	}
}
