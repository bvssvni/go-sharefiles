package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Displays a list of files in the "shared" directory.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>")
	filepath.Walk("./shared", 
		func (path string, f os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if (f.IsDir()) {
				return nil
			}
			printFile(w, path)
			return nil
		})
	fmt.Fprintf(w, "</body></html>")
}

// Print out a link to file in web browser.
func printFile(w http.ResponseWriter, file string) {
	fmt.Fprintf(w, "<a target=\"_blank\" href=\"")
	fmt.Fprintf(w, "%s\n", file)
	fmt.Fprintf(w, "\">")
	fmt.Fprintf(w, "%s\n", file)
	fmt.Fprintf(w, "</a><br />")
}

// Reads from shared directory and displays the file in browser.
func sharedHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}
	fmt.Fprintf(w, "%s", data)
}

// Starts web server.
func main() {
	http.HandleFunc("/shared/", sharedHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


