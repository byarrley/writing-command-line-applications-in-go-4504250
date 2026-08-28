package main

import "os"

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
		   inFile: path to input file (default: stdin)
		     * Ensure that the file path is valid
	       * File is non-empty text file
	*/

	text := "Go"
	width := 6
	Banner(os.Stdout, text, width)
}
