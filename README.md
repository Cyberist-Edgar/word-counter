## word counter
Usage: `wc [OPTION]... [FILE]...`
```
go run main.go test README.md
3 6 27 test
15 87 495 README.md
```

Print newline, word, and byte counts for each FILE. A word is a non-zero-length sequence of
characters delimited by white space.

The options below may be used to select which counts are printed, always in
the following order: newline, word, byte

`-c` print the byte counts

`-l` print the newline counts

`-w` print the word counts

`--help` display this help and exit

`-v` output version information and exit
