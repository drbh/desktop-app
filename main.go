package main

import (
	"fmt"
	"regexp"

	"flag"

	"io"
	"log"
	"os"
)

const (
	s = `
package main

import "github.com/zserge/webview"

func main() {
	webview.Open("REPLACE_TITLE", "REPLACE_URL", WIDTH, HEIGHT, true)
}
`
)

func main() {

	title_pointer := flag.String("title", "My Desktop App", "The title of the web ui window")
	url_pointer := flag.String("url",
		"https://medium.com/@david.richard.holtz/how-facebook-is-going-to-change-crypto-forever-155498c22b0",
		"The webapp's URL")

	height_pointer := flag.String("height", "900", "The title of the web ui window")
	width_pointer := flag.String("width", "900", "The webapp's URL")

	flag.Parse()

	r := regexp.MustCompile("HEIGHT")
	val1 := r.ReplaceAllString(s, *height_pointer)

	r = regexp.MustCompile("WIDTH")
	val2 := r.ReplaceAllString(val1, *width_pointer)

	r = regexp.MustCompile("REPLACE_TITLE")
	val3 := r.ReplaceAllString(val2, *title_pointer)

	r = regexp.MustCompile("REPLACE_URL")
	final := r.ReplaceAllString(val3, *url_pointer)
	fmt.Println(final)

	err := WriteToFile("app.go", final)
	if err != nil {
		log.Fatal(err)
	}

}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
