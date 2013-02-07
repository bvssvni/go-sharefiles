/*
	ShareFiles - A program for sharing files with other machines through the browser. 
	https://github.com/bvssvni/go-sharefiles
	BSD license.
	by Sven Nilsen, 2013
	http://www.cutoutpro.com

	Version: 0.000 in angular degrees version notation
	http://isprogrammingeasy.blogspot.no/2012/08/angular-degrees-versioning-notation.html
*/
/*
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
The views and conclusions contained in the software and documentation are those
of the authors and should not be interpreted as representing official policies,
either expressed or implied, of the FreeBSD Project.
*/

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


