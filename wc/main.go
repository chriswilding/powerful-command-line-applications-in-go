package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *bytes, *lines))
}

func count(r io.Reader, countBytes, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	var wc int
	for scanner.Scan() {
		wc++
	}

	return wc
}
