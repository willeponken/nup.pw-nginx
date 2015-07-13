package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"regexp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(`Copyright (C) 2015 oniichaNj
This program is free software:
You can redistribute it and/or modify it under the terms of the GNU General Public License
as published by the Free Software Foundation, version. This program is distributed in the
hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with this program.
If not, see <http://www.gnu.org/licenses/>.`)
		fmt.Printf("\nUsage: %s file1 file2 .. fileN\n", os.Args[0])
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		fmt.Printf("%s ", file)
		err := Upload("https://nup.pw", file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Upload(url, file string) (err error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open(file)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("file", path.Base(file))
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	fmt.Println("uploading.")
	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	//yo i heard you liked spaghetti
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	re := regexp.MustCompile(`(https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	fmt.Println(re.FindAllString(buf.String(), -1)[1])
	return

}
