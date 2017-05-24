Stdout performance
==================

I have been just using `fmt.Println()` or `file.Write()` to output things. But
it can be inefficient to write many small data if [File][file] is unbuffered.
The document says nothing. Should I wrap a file with [bufio.Writer][bufio]?

[file]: https://golang.org/pkg/os/#File
[bufio]: https://golang.org/pkg/bufio/#Writer

## The benchmark

The program measures the time required to write many strings to Stdout with
or without explicit buffering. The output should be redirected to /dev/null
since writing to a physical drive involves extra cost that is irrelevant for
this benchmark.

    go build
    ./12-ioperformance -u 10000000 > /dev/null
    ./12-ioperformance -b 10000000 > /dev/null

On my machine, the program ran 24 times faster when buffering is on:

| Buffered | Time (sec) |
|----------|------------|
| No       | 2.4        |
| Yes      | 0.098      |

So yes, I should explicitly turn on buffering when writing many small data.
