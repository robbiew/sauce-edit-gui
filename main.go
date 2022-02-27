// Demo code for the Form primitive.
package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/rivo/tview"
	sauce "github.com/ActiveState/go-ansi"
)


var (
	sGroup    string
	sAuthor   string
	sTitle    string
	artPath   string
)
func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pathPtr := flag.String("path", "", "path to ANSI file with SAUCE")
	required := []string{"path"}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "\nerror: -%s argumant is reuired! Exiting...\r\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
	artPath = dir + "/" + *pathPtr
}

func main() {
		// let's check the file for a valid SAUCE record
		record := sauce.GetSauce(artPath)

			// abort if we don't find SAUCE
	if string(record.Sauceinf.ID[:]) == sauce.SauceID {

		fmt.Printf("----------------------------------------\n" )
	} else {
		fmt.Println("NO SAUCE RECORD FOUND!")
		os.Exit(0)
	}

	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Title", fmt.Sprintf("%v",record.Sauceinf.Title), 35, nil, nil).
		AddInputField("Author", record.Sauceinf.Author, 20, nil, nil).
		AddInputField("Group", "", 20, nil, nil).
		AddInputField("Comments","",64, nil, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}