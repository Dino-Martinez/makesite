package main

import (
	"fmt"
	"io/ioutil"
	"html/template"
	"os"
	"flag"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	dir := flag.String("dir", ".", "The directory from which to parse all txt files")
	//filePath := flag.String("file", "first-post.txt", "The name of the file to generate HTML of, including file extension.")
	flag.Parse()

	// fileContents, err := ioutil.ReadFile(*filePath)

	// if err != nil {
	// 	// A common use of `panic` is to abort if a function returns an error
	// 	// value that we don’t know how to (or want to) handle. This example
	// 	// panics if we get an unexpected error when creating a new file.
	// 	panic(err)
	// }

	files, err := ioutil.ReadDir(*dir)

	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	
		// fileName := string(*filePath)[:len(*filePath) - 4]
		filePath := file.Name()
		fileName := filePath[:len(filePath) - 4]
		fileExtension := filePath[len(filePath) - 4:]

		if fileExtension == ".txt" {
			fileContents, err := ioutil.ReadFile(filePath)

			fmt.Println(string(fileContents), fileName)

			page := Page{
				TextFilePath: filePath,
				TextFileName: fileName,
				HTMLPagePath: fileName + ".html",
				Content:      string(fileContents),
			}

			// Create a new template in memory named "template.tmpl".
			// When the template is executed, it will parse template.tmpl,
			// looking for {{ }} where we can inject content.
			t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

			// Create a new, blank HTML file.
			newFile, err := os.Create(page.HTMLPagePath)
			if err != nil {
						panic(err)
			}

			// Executing the template injects the Page instance's data,
			// allowing us to render the content of our text file.
			// Furthermore, upon execution, the rendered template will be
			// saved inside the new file we created earlier.
			t.Execute(newFile, page)
		}
	}
}
