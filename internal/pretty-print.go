package internal

import (
	"log"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/hoehwa/gopkg/bubbletea/listfancy"
)

// Read and return contents of the file at the specific absoulte path.
func readFileAt(fileAbsPath string) string {
	readStream, err := os.ReadFile(fileAbsPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(readStream)
}

// Return a pretty-printed filename as an URL format.
func fmtFileNameAsURL(subPath string, fname string) string {
	basename := Baseurl + subPath + "/"
	location := strings.TrimRight(fname, ".")

	return basename + location
}

// Return titles string array and descriptions string array respectively.
func getTitlesAndDescs(srcPath string, subPath string) ([]string, []string) {
	files, err := os.ReadDir(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 0, len(files))
	descs := make([]string, 0, len(files))

	for _, file := range files {
		titles = append(titles, file.Name())
		descs = append(descs, fmtFileNameAsURL(subPath, file.Name()))
	}

	return titles, descs
}

// Print content of selected item with highlighted text.
func PrettyPrint(subPath string) {
	srcPath := BaseDir + subPath

	titles, descs := getTitlesAndDescs(srcPath, subPath)
	var selection *string

	// Class to generate titles and descriptions into the bubble-tea client.
	// It also save user a selection and pass it to another process to show contents for a selected file.
	callout := listfancy.NewCallout(listfancy.Callout{
		Titles:    titles,
		Descs:     descs,
		Selection: selection,
	})

	listfancy.InitCallout(*callout)

	selectedFile := *callout.Selection
	contents := readFileAt(srcPath + "/" + selectedFile)

	if err := quick.Highlight(os.Stdout, contents, selectedFile, "terminal", "monokai"); err != nil {
		log.Fatal(err)
	}
}
