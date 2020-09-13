package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	bytes       bool
	lines       bool
	words       bool
	showVersion bool
)

func init() {
	flag.BoolVar(&bytes, "c", false, "print the byte counts")
	flag.BoolVar(&lines, "l", false, "print the newline counts")
	flag.BoolVar(&words, "w", false, "print the word counts")
	flag.BoolVar(&showVersion, "v", false, "output version information and exit")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: wc [OPTION]... [FILE]...")
	fmt.Fprintln(os.Stderr, "Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified. A word is a noe-zero-length sequence of characters delimited by white space")
	fmt.Fprintln(os.Stderr, "Options: ")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if showVersion {
		fmt.Fprintln(os.Stderr, "wc Version 1.0")
		fmt.Fprintln(os.Stderr, "Copyright 2020 Edgar, visit https://git.io/JU8sQ for more detail")
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Must specify [FILE]")
		os.Exit(0)
	}

	// 所需要进行处理的file文件
	for _, value := range args {
		l, w, b := readFile(value)
		if !(lines || words || bytes){
			lines, words, bytes = true, true, true
		}
		if lines{
			fmt.Printf("%v ", l)
		}
		if words{
			fmt.Printf("%v ",w)
		}
		if bytes{
			fmt.Printf("%v ",b)
		}
		fmt.Println(value)
	}
}

func readFile(filename string) (lines, words, bytes int) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "wc: %s: No such file\n", filename)
		os.Exit(2)
	}
	if info.IsDir() {
		fmt.Fprintf(os.Stderr, "wc: %s: Is a directory\n", filename)
		os.Exit(2)
	}
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		if err == os.ErrPermission {
			fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", filename)
		}
		os.Exit(2)
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// 说明读完了
			if err == io.EOF {
				lines++
				b, w := dealLine(line)
				bytes += b
				words += w
				return
			}
			fmt.Fprintf(os.Stderr, "wc: %s: Read file error\n", filename)
			os.Exit(2)
		}
		lines++
		b, w := dealLine(line)
		bytes += b
		words += w

	}
	return
}

func dealLine(line string) (bytes, words int) {
	line =  strings.Replace(line, "\n", "", -1)
	line =  strings.Replace(line, "\r", "", -1)

	bytes += len(line)
	for _, value := range strings.SplitN(line, " ", -1) {
		if len(value) > 0 && value != "\n"{
			words++
		}
	}
	return
}
