package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
Instructions:

- Pass text to print as an argument
  - If not argument - read from stdin

- Use -width to specify width
  - Width should be bigger than 0 and less than 250
  - Default to 80

- Use -out to specify output file
  - Default to stdout
*/

func main() {

	/* Flags
	   -width (bannerWidth): Banner width, 0 < width < 250 (default: 80)
	    * Ensure -width is an int
	    * Ensure that it is between 0 and 250
	   -out (outFile): path to output file (default: stdout)
	    * Ensure that the file can be opened
	*/

	/* Args
	   text: text to print (default: stdin)
	     * This argument doesn't require file ops
			 * Unclear from video/comments what to do with multiline stdin, so print a banner for each line
			 * For multiple args, assume that each is a separate input with its own banner (consistent with stdin treatment)
	*/

	//TODO: move banner.go, args to separate packages
	var out outStream

	flag.Var(&out, "out", "Path to output file (default: stdout)")
	flag.Parse()

	text := readInput()

	width := 6

	for _, line := range text {
		Banner(out.stream, line, width)
	}

}

func readInput() []string {
	var t []string
	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			t = append(t, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	} else {
		t = flag.Args()
	}
	return t
}
