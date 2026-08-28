package main

import (
	"bufio"
	"flag"
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
	*/
	flag.Parse()

	text := readInput()
	width := 6
	Banner(os.Stdout, text, width)
}

func readInput() string {
	var t string
	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			t = scanner.Text()
		}
	} else {
		t = flag.Arg(0)
	}
	return t
}
