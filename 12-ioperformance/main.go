package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	opts, err := parseOptions()
	if err != nil {
		return err
	}
	var writer io.Writer = os.Stdout
	if opts.buffered {
		writer = bufio.NewWriter(writer)
	}
	duration, err := benchmark(writer, opts.benchSize)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Completed in", duration)
	return nil
}

type options struct {
	buffered  bool
	benchSize int
}

func parseOptions() (*options, error) {
	opts := options{
		buffered:  false,
		benchSize: 10000000,
	}
	if len(os.Args) > 3 {
		return nil, fmt.Errorf("unrecognized options: %s", os.Args[3:])
	}
	if len(os.Args) > 2 {
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return nil, err
		}
		opts.benchSize = n
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-u":
			opts.buffered = false
		case "-b":
			opts.buffered = true
		default:
			return nil, fmt.Errorf("unrecognized option: %s", os.Args[1])
		}
	}
	return &opts, nil
}

func benchmark(writer io.Writer, benchSize int) (time.Duration, error) {
	data := []byte("Lorem Ipsum\n")
	start := time.Now()
	for i := 0; i < benchSize; i++ {
		if _, err := writer.Write(data); err != nil {
			return 0, err
		}
	}
	if buf, buffered := writer.(*bufio.Writer); buffered {
		buf.Flush()
	}
	return time.Since(start), nil
}
